BEGIN TRANSACTION;
CREATE TABLE "users" ("id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, "email" varchar DEFAULT '' NOT NULL, "encrypted_password" varchar DEFAULT '' NOT NULL, "reset_password_token" varchar, "reset_password_sent_at" datetime, "remember_created_at" datetime, "sign_in_count" integer DEFAULT 0 NOT NULL, "current_sign_in_at" datetime, "last_sign_in_at" datetime, "current_sign_in_ip" varchar, "last_sign_in_ip" varchar, "created_at" datetime NOT NULL, "updated_at" datetime NOT NULL);
INSERT INTO `users` VALUES (1,'dchen15@simons-rock.edu','$2a$10$RYE4bcb2QcHL0R0lPJtcCe1RY2UxQt.oxFPygNGhD1ti2jgF2CaAu',NULL,NULL,NULL,2,'2016-02-08 01:49:00.901866','2016-02-08 01:38:38.052188','127.0.0.1','127.0.0.1','2016-02-08 01:38:38.025816','2016-02-08 01:49:00.902607');
CREATE TABLE "schema_migrations" ("version" varchar NOT NULL);
INSERT INTO `schema_migrations` VALUES ('20160206184821');
INSERT INTO `schema_migrations` VALUES ('20160208013215');
INSERT INTO `schema_migrations` VALUES ('20160209161915');
CREATE TABLE "responses" ("id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, "commenter" varchar, "body" text, "challenge_id" integer, "created_at" datetime NOT NULL, "updated_at" datetime NOT NULL);
INSERT INTO `responses` VALUES (3,'dchen15','#inlcude,Noob.h>
int return',2,'2016-02-12 23:42:38.349686','2016-02-12 23:42:38.349686');
CREATE TABLE "challenges" ("id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, "title" varchar, "details" text, "created_at" datetime NOT NULL, "updated_at" datetime NOT NULL);
INSERT INTO `challenges` VALUES (1,'Find String in Substring','These are the details for the String in Substring.','2016-02-06 19:38:56.368226','2016-02-06 19:38:56.368226');
INSERT INTO `challenges` VALUES (2,'Somethign','Sdfsajlkaf
','2016-02-06 21:04:35.113452','2016-02-06 21:04:35.113452');
INSERT INTO `challenges` VALUES (3,'New CHalleng','Print WHldf DF','2016-02-12 23:43:17.932819','2016-02-12 23:43:17.932819');
INSERT INTO `challenges` VALUES (4,'This is a Sample Article','Quot etiam numquam cum cu, nec elit fierent oporteat no. In facer accusata cum, nec ad ipsum legendos necessitatibus. Malis lucilius et vel, pro causae scripserit an. Sed ne choro ignota consectetuer. Modo munere populo has ut, ne vix quot contentiones.\r\n\r\nGraeco accusam eu per. Tation rationibus cum an, ut usu ferri posse iudico, solet legendos te sit. Splendide theophrastus id pro, ad fugit iriure volumus mei. Tation neglegentur ad mei, civibus mandamus honestatis te mei, unum paulo utamur pro at. Affert maluisset ullamcorper pri eu. Inermis appetere id vix, ea sed prima aliquid impedit.','2016-03-17 01:11:21.154651','2016-03-17 01:11:21.154651');
INSERT INTO `challenges` VALUES (5,'This is a Sample Article','Quot etiam numquam cum cu, nec elit fierent oporteat no. In facer accusata cum, nec ad ipsum legendos necessitatibus. Malis lucilius et vel, pro causae scripserit an. Sed ne choro ignota consectetuer. Modo munere populo has ut, ne vix quot contentiones.\r\n\r\nGraeco accusam eu per. Tation rationibus cum an, ut usu ferri posse iudico, solet legendos te sit. Splendide theophrastus id pro, ad fugit iriure volumus mei. Tation neglegentur ad mei, civibus mandamus honestatis te mei, unum paulo utamur pro at. Affert maluisset ullamcorper pri eu. Inermis appetere id vix, ea sed prima aliquid impedit.','2016-03-17 01:11:21.154651','2016-03-17 01:11:21.154651');
INSERT INTO `challenges` VALUES (6,'This is a Sample Article','Quot etiam numquam cum cu, nec elit fierent oporteat no. In facer accusata cum, nec ad ipsum legendos necessitatibus. Malis lucilius et vel, pro causae scripserit an. Sed ne choro ignota consectetuer. Modo munere populo has ut, ne vix quot contentiones.\r\n\r\nGraeco accusam eu per. Tation rationibus cum an, ut usu ferri posse iudico, solet legendos te sit. Splendide theophrastus id pro, ad fugit iriure volumus mei. Tation neglegentur ad mei, civibus mandamus honestatis te mei, unum paulo utamur pro at. Affert maluisset ullamcorper pri eu. Inermis appetere id vix, ea sed prima aliquid impedit.','2016-03-17 01:11:21.154651','2016-03-17 01:11:21.154651');
CREATE UNIQUE INDEX "unique_schema_migrations" ON "schema_migrations" ("version");
CREATE UNIQUE INDEX "index_users_on_reset_password_token" ON "users" ("reset_password_token");
CREATE UNIQUE INDEX "index_users_on_email" ON "users" ("email");
CREATE INDEX "index_responses_on_challenge_id" ON "responses" ("challenge_id");
COMMIT;
