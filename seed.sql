-- DROP DATABASE ads; 
-- CREATE DATABASE ads;

DROP TABLE IF EXISTS ads;
DROP TABLE IF EXISTS ad_media;
DROP TABLE IF EXISTS impressions;
DROP TABLE IF EXISTS events;

CREATE TABLE IF NOT EXISTS ads (
  id SERIAL PRIMARY KEY,
  title varchar(40),
  description varchar(160),
  priority integer,
  zone_id integer
);

CREATE TABLE IF NOT EXISTS ad_media (
  id SERIAL PRIMARY KEY, 
  ad_id integer,
  media_type varchar(20),
  height integer, 
  width integer,
  delivery varchar(20),
  url varchar(100)
);

CREATE TABLE IF NOT EXISTS impressions (
  id SERIAL PRIMARY KEY,
  ad_id integer, 
  time timestamp without time zone default (now() at time zone 'utc')
);

CREATE TABLE IF NOT EXISTS events (
  id SERIAL PRIMARY KEY,
  ad_id integer, 
  action varchar(60),
  time timestamp without time zone default (now() at time zone 'utc')
);

INSERT INTO ads (title, description, priority, zone_id) VALUES ('title1', 'description1', 8, 1);
INSERT INTO ads (title, description, priority, zone_id) VALUES ('title2', 'description2', 9, 1);
INSERT INTO ads (title, description, priority, zone_id) VALUES ('title3', 'description3', 10, 1);
INSERT INTO ads (title, description, priority, zone_id) VALUES ('title4', 'description4', 10, 2);
INSERT INTO ads (title, description, priority, zone_id) VALUES ('title5', 'description5', 9, 3);

INSERT INTO ad_media(ad_id, media_type, height, width, delivery, url) VALUES (1, 'video/mp4', 720, 405, 'progressive', 'http://ad-videos.host.com/videos/video1.mp4');
INSERT INTO ad_media(ad_id, media_type, height, width, delivery, url) VALUES (1, 'video/mp4', 720, 405, 'progressive', 'http://ad-videos.host.com/videos/video2.mp4'); 

INSERT INTO ad_media(ad_id, media_type, height, width, delivery, url) VALUES (2, 'video/mp4', 720, 405, 'progressive', 'http://ad-videos.host.com/videos/video3.mp4'); 
INSERT INTO ad_media(ad_id, media_type, height, width, delivery, url) VALUES (2, 'video/mp4', 720, 405, 'progressive', 'http://ad-videos.host.com/videos/video4.mp4');

INSERT INTO ad_media(ad_id, media_type, height, width, delivery, url) VALUES (3, 'video/mp4', 720, 405, 'progressive', 'http://ad-videos.host.com/videos/video5.mp4');
INSERT INTO ad_media(ad_id, media_type, height, width, delivery, url) VALUES (3, 'video/mp4', 720, 405, 'progressive', 'http://ad-videos.host.com/videos/video6.mp4');

INSERT INTO ad_media(ad_id, media_type, height, width, delivery, url) VALUES (4, 'video/mp4', 720, 405, 'progressive', 'http://ad-videos.host.com/videos/video7.mp4');
INSERT INTO ad_media(ad_id, media_type, height, width, delivery, url) VALUES (4, 'video/mp4', 720, 405, 'progressive', 'http://ad-videos.host.com/videos/video8.mp4');

INSERT INTO ad_media(ad_id, media_type, height, width, delivery, url) VALUES (5, 'video/mp4', 720, 405, 'progressive', 'http://ad-videos.host.com/videos/video9.mp4');
INSERT INTO ad_media(ad_id, media_type, height, width, delivery, url) VALUES (5, 'video/mp4', 720, 405, 'progressive', 'http://ad-videos.host.com/videos/video10.mp4');

