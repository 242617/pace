package storage

const queryGetUser = `

	select face_id
	from public.user
	where phone = $1;

`
