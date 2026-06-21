CREATE TABLE public.sessions (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	summary varchar(256) NOT NULL,
	"location" varchar(64) NOT NULL,
	start_time timestamptz NOT NULL,
	CONSTRAINT sessions_pkey PRIMARY KEY (id)
);

CREATE INDEX idx_session_start_time ON public.sessions USING btree (start_time);
