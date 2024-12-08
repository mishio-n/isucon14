package main

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/oklog/ulid/v2"
)

/**
 * [x] ランダムな椅子の選択　から 最も近い椅子を選択に変更した
 * [ ] 椅子の状態をより詳細に管理し、例えば「移動中」「待機中」などの状態を追加する
 * [ ] ライドリクエストの優先度を考慮し、例えばVIPユーザーのリクエストを優先するなどの方法があります
 */
func internalGetMatching(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ride := &Ride{}
	if err := db.GetContext(ctx, ride, `SELECT * FROM rides WHERE chair_id IS NULL ORDER BY created_at LIMIT 1`); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	// 最も近い椅子を選択する
	matched := &Chair{}
	if err := db.GetContext(ctx, matched, `
        SELECT *, 
            (6371 * acos(cos(radians(?)) * cos(radians(latitude)) * cos(radians(longitude) - radians(?)) + sin(radians(?)) * sin(radians(latitude)))) AS distance 
        FROM chairs 
        WHERE is_active = TRUE 
        ORDER BY distance 
        LIMIT 1`, ride.PickupLatitude, ride.PickupLongitude, ride.PickupLatitude); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if _, err := db.ExecContext(ctx, `UPDATE rides SET chair_id = ? WHERE id = ?`, matched.ID, ride.ID); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if _, err := db.ExecContext(ctx, `INSERT INTO ride_statuses (id, ride_id, status) VALUES (?, ?, ?)`, ulid.Make().String(), ride.ID, "MATCHING"); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
