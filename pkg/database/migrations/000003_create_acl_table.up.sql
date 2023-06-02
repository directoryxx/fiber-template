CREATE TABLE IF NOT EXISTS acl (
    p_type varchar(256) not null default '',
    v0 		varchar(256) not null default '',
    v1 		varchar(256) not null default '',
    v2 		varchar(256) not null default '',
    v3 		varchar(256) not null default '',
    v4 		varchar(256) not null default '',
    v5 		varchar(256) not null default ''
);

CREATE INDEX IF NOT EXISTS idx_acl_v5 ON acl(v5);
CREATE INDEX IF NOT EXISTS idx_acl_v4 ON acl(v4);
CREATE INDEX IF NOT EXISTS idx_acl_v3 ON acl(v3);
CREATE INDEX IF NOT EXISTS idx_acl_v2 ON acl(v2);
CREATE INDEX IF NOT EXISTS idx_acl_v1 ON acl(v1);
CREATE INDEX IF NOT EXISTS idx_acl_v0 ON acl(v0);
CREATE INDEX IF NOT EXISTS idx_acl_p_type ON acl(p_type);