package storage

const (
	queryGetUserByPhone = `
		select phone, person_id, name, alias
		from public.user
		where phone = $1;
	`

	queryGetUserByPersonID = `
		select phone, person_id, name, alias
		from public.user
		where person_id = $1;
	`

	queryCreateUser = `
		insert into public.user (phone)
		values ($1) returning id;
	`

	queryUpdateUserName = `
		update public.user
		set name = $1
		where phone = $2;
	`

	queryUpdateUserAlias = `
		update public.user
		set alias = $1
		where phone = $2;
	`

	queryUpdateUserPersonID = `
		update public.user
		set person_id = $1
		where phone = $2;
	`
)
