CREATE TABLE account (
  id INT GENERATED ALWAYS AS IDENTITY,
  email varchar(256) NOT NULL,
  password varchar(256) NOT NULL,
  PRIMARY KEY(id)
);

INSERT INTO account(email,password) VALUES ('admin@admin.ru','d033e22ae348aeb5660fc2140aec35850c4da997');

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
