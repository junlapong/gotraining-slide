CREATE TABLE posts (
       id INT NOT NULL AUTO_INCREMENT,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
       title TEXT NOT NULL,
       body TEXT NOT NULL,
       PRIMARY KEY (id)
);
