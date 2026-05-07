CREATE TABLE IF NOT EXISTS projects (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL DEFAULT '',
    priority VARCHAR(50) DEFAULT '',
    stage VARCHAR(100) DEFAULT '',
    start_date VARCHAR(50) DEFAULT '',
    deadline VARCHAR(50) DEFAULT '',
    cost VARCHAR(50) DEFAULT '',
    paid VARCHAR(50) DEFAULT '',
    remainder VARCHAR(50) DEFAULT '',
    other_income VARCHAR(50) DEFAULT '',
    revenue VARCHAR(50) DEFAULT '',
    expense VARCHAR(50) DEFAULT '',
    profit VARCHAR(50) DEFAULT '',
    description TEXT DEFAULT '',
    customer VARCHAR(255) DEFAULT '',
    executor VARCHAR(255) DEFAULT '',
    tax_rate VARCHAR(50) DEFAULT '',
    account VARCHAR(100) DEFAULT '',
    end_date VARCHAR(50) DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_projects_user_id ON projects(user_id);
