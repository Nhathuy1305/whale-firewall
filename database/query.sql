-- name: AddContainer :exec
INSERT INTO containers(id, name)
VALUES (?, ?);

-- name: AddContainerAddr :exec
INSERT INTO addrs(addr, container_id)
VALUES (?, ?);

-- name: AddContainerAlias :exec
INSERT INTO container_aliases(container_id, container_alias)
VALUES (?, ?);

-- name: AddEstContainer :exec
INSERT INTO est_containers(src_container_id, dst_container_id)
VALUES (?, ?);

-- name: AddWaitingContainerRule :exec
INSERT INTO waiting_container_rules(src_container_id, dst_container_name, rule)
VALUES (?, ?, ?)
ON CONFLICT(src_container_id, dst_container_name, rule) DO NOTHING;

// TODO
