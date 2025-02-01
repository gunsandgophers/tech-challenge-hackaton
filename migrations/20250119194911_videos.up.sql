CREATE TABLE videos
(
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'AWAITING',
    mime_type VARCHAR(50) NOT NULL,
    filename VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT pk_videos_id PRIMARY KEY (id),
    CONSTRAINT status CHECK (status IN ('AWAITING', 'CANCELED', 'FINISHED'))
);
