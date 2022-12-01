CREATE TABLE cakes (
    id int NOT NULL AUTO_INCREMENT,
    title varchar(255) NOT NULL UNIQUE,
    description varchar(255),
    rating float,
    image varchar(255),
  	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
  	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

INSERT IGNORE INTO cakes 
  (title, description, rating, image)
VALUES 
  ('[Seeder Data] Chocolate cheesecakes', 'A cheesecake made of cheesecakes', 3, 'https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg');
