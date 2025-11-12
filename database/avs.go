package database

import "fmt"

// SaveAvsAddr saves the given avs address into the database.
// TODO: implement a full-fledged function that saves all AVS information
func (db *Db) SaveAvsAddr(avsAddr string) error {
	stmt := `
INSERT INTO avs (avs_addr)
VALUES ($1)
ON CONFLICT (avs_addr) DO NOTHING;`
	_, err := db.SQL.Exec(stmt, avsAddr)
	if err != nil {
		return fmt.Errorf("error while saving avs address: %s", err)
	}
	return nil
}

// SaveChainIdToAvsAddr saves the given chain id to avs address mapping into the database.
func (db *Db) SaveChainIdToAvsAddr(chainID, avsAddr string) error {
	stmt := `
INSERT INTO chain_id_to_avs_addr (chain_id, avs_addr)
VALUES ($1, $2)
ON CONFLICT (chain_id) DO UPDATE
SET avs_addr = EXCLUDED.avs_addr;`
	_, err := db.SQL.Exec(stmt, chainID, avsAddr)
	if err != nil {
		return fmt.Errorf("error while saving chain id to avs address: %s", err)
	}
	return nil
}

// GetAllAvsAddrs returns all AVS addresses from the database.
func (db *Db) GetAllAvsAddrs() ([]string, error) {
	stmt := `SELECT avs_addr FROM avs ORDER BY avs_addr;`

	rows, err := db.SQL.Query(stmt)
	if err != nil {
		return nil, fmt.Errorf("error while querying AVS addresses: %w", err)
	}
	defer rows.Close()

	var avsAddrs []string
	for rows.Next() {
		var addr string
		if err := rows.Scan(&addr); err != nil {
			return nil, fmt.Errorf("error while scanning AVS address: %w", err)
		}
		avsAddrs = append(avsAddrs, addr)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return avsAddrs, nil
}
