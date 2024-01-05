CREATE TABLE account (
  id INT GENERATED ALWAYS AS IDENTITY,
  email varchar(256) NOT NULL,
  username varchar(256) NOT NULL,
  password varchar(256) NOT NULL,
  totp_secret bytea,
  PRIMARY KEY(id)
);

INSERT INTO account(email,password,totp_secret) VALUES ('admin@admin.ru','f865b53623b121fd34ee5426c792e5c33af8c227',E'\\xAAAAAAAAAAAAAAAAAAAA');

CREATE TABLE account_role (
  id INT GENERATED ALWAYS AS IDENTITY,
  account_id INT NOT NULL,
  role varchar(64) NOT NULL,
  PRIMARY KEY(id),
  CONSTRAINT fk_account_role_account_id
    FOREIGN KEY(account_id) 
	    REFERENCES account(id)
);

INSERT INTO account_role(account_id,role) VALUES (1,'admin'), (1,'user');

CREATE TABLE session (
  id INT GENERATED ALWAYS AS IDENTITY,
  account_id INT NOT NULL,
  access_token varchar(64) NOT NULL,
  refresh_token varchar(64) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE,
  updated_at TIMESTAMP WITH TIME ZONE,
  PRIMARY KEY(id),
  CONSTRAINT fk_session_account_id
    FOREIGN KEY(account_id) 
	    REFERENCES account(id)
);

CREATE TABLE log (
  id INT GENERATED ALWAYS AS IDENTITY,
  account_id INT,
  session_id INT,
  level varchar(5) NOT NULL,
  message TEXT,
  data TEXT,
  created_at TIMESTAMP WITH TIME ZONE,
  PRIMARY KEY(id),
  CONSTRAINT fk_log_account_id
    FOREIGN KEY(account_id) 
	    REFERENCES account(id),
  CONSTRAINT fk_log_session_id
    FOREIGN KEY(session_id) 
	    REFERENCES session(id)
);

CREATE TABLE login_attempt (
  id INT GENERATED ALWAYS AS IDENTITY,
  ip TEXT,
  count INT,
  last_failed TIMESTAMP WITH TIME ZONE,
  PRIMARY KEY(id)
);

CREATE TABLE remote_account (
  id INT GENERATED ALWAYS AS IDENTITY,
  account_id INT NOT NULL,
  remote_id text NOT NULL,
  remote_email varchar(256) NOT NULL,
  platform varchar(64) NOT NULL,
  PRIMARY KEY(id),
  CONSTRAINT fk_remote_account_account_id
    FOREIGN KEY(account_id) 
	    REFERENCES account(id)
);
