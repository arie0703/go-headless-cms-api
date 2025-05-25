-- users テーブル
CREATE TABLE users (
  id CHAR(36) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password_hash TEXT NOT NULL,
  role ENUM('admin', 'editor', 'viewer') NOT NULL DEFAULT 'editor',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- content_types テーブル
CREATE TABLE content_types (
  id CHAR(36) PRIMARY KEY,
  name VARCHAR(100) NOT NULL UNIQUE,
  display_name VARCHAR(100),
  description TEXT,
  created_by CHAR(36),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- fields テーブル
CREATE TABLE fields (
  id CHAR(36) PRIMARY KEY,
  content_type_id CHAR(36) NOT NULL,
  name VARCHAR(100) NOT NULL,
  label VARCHAR(100),
  type ENUM('text', 'number', 'boolean', 'date', 'reference', 'richtext', 'image') NOT NULL,
  is_required BOOLEAN DEFAULT FALSE,
  `order` INT DEFAULT 0,
  settings JSON,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 初期ユーザー & コンテンツタイプ追加
INSERT INTO users (id, name, email, password_hash, role)
VALUES (UUID(), 'Admin', 'admin@example.com', 'your_password_hash', 'admin');

INSERT INTO content_types (id, name, display_name, description)
VALUES (UUID(), 'blog', 'ブログ', 'ブログ投稿');

INSERT INTO fields (id, content_type_id, name, label, type, is_required, `order`)
VALUES 
(UUID(), (SELECT id FROM content_types WHERE name='blog'), 'title', 'タイトル', 'text', TRUE, 1),
(UUID(), (SELECT id FROM content_types WHERE name='blog'), 'body', '本文', 'richtext', TRUE, 2);
