BEGIN TRANSACTION;
DROP TABLE IF EXISTS "roles";
CREATE TABLE IF NOT EXISTS "roles" (
	"id"	integer,
	"name"	text NOT NULL UNIQUE,
	"description"	text NOT NULL UNIQUE,
	"active"	numeric DEFAULT true,
	"app_id"	integer,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_apps_roles" FOREIGN KEY("app_id") REFERENCES "apps"("id")
);
DROP TABLE IF EXISTS "user_roles";
CREATE TABLE IF NOT EXISTS "user_roles" (
	"user_id"	integer,
	"role_id"	integer,
	CONSTRAINT "fk_user_roles_user" FOREIGN KEY("user_id") REFERENCES "users"("id") ON DELETE SET NULL ON UPDATE CASCADE,
	CONSTRAINT "fk_user_roles_role" FOREIGN KEY("role_id") REFERENCES "roles"("id") ON DELETE SET NULL ON UPDATE CASCADE,
	PRIMARY KEY("user_id","role_id")
);
DROP TABLE IF EXISTS "user_roles constraint:_on_update:_cascad_e,_on_delete:_se_t _nulls";
CREATE TABLE IF NOT EXISTS "user_roles constraint:_on_update:_cascad_e,_on_delete:_se_t _nulls" (
	"role_id"	integer,
	"user_id"	integer,
	PRIMARY KEY("role_id","user_id"),
	CONSTRAINT "fk_user_roles constraint:_on_update:_cascad_e,_on_deleteae4dc996" FOREIGN KEY("user_id") REFERENCES "users"("id"),
	CONSTRAINT "fk_user_roles constraint:_on_update:_cascad_e,_on_delete0b805828" FOREIGN KEY("role_id") REFERENCES "roles"("id")
);
DROP TABLE IF EXISTS "pages";
CREATE TABLE IF NOT EXISTS "pages" (
	"id"	integer,
	"name"	text UNIQUE,
	"active"	numeric DEFAULT true,
	"description"	text,
	PRIMARY KEY("id")
);
DROP TABLE IF EXISTS "features";
CREATE TABLE IF NOT EXISTS "features" (
	"id"	integer,
	"name"	text UNIQUE,
	"description"	text,
	"active"	numeric DEFAULT true,
	"role_id"	integer,
	CONSTRAINT "fk_roles_features" FOREIGN KEY("role_id") REFERENCES "roles"("id") ON DELETE SET NULL ON UPDATE CASCADE,
	PRIMARY KEY("id")
);
DROP TABLE IF EXISTS "session_data";
CREATE TABLE IF NOT EXISTS "session_data" (
	"token"	text,
	"time_stamp"	datetime DEFAULT current_timestamp
);
DROP TABLE IF EXISTS "site_data";
CREATE TABLE IF NOT EXISTS "site_data" (
	"id"	integer,
	"remote_add"	varchar(128),
	"accessed_route"	varchar(300),
	"method"	varchar(10),
	"response_time"	real,
	"response_status"	integer,
	PRIMARY KEY("id")
);
DROP TABLE IF EXISTS "apps";
CREATE TABLE IF NOT EXISTS "apps" (
	"id"	integer,
	"name"	text UNIQUE,
	"uuid"	uuid,
	"active"	numeric DEFAULT true,
	"description"	text,
	PRIMARY KEY("id")
);
DROP TABLE IF EXISTS "users";
CREATE TABLE IF NOT EXISTS "users" (
	"id"	integer,
	"uuid"	uuid,
	"email"	text UNIQUE,
	"password"	text,
	"date_registered"	datetime DEFAULT current_timestamp,
	"disabled"	numeric DEFAULT false,
	PRIMARY KEY("id")
);
DROP TABLE IF EXISTS "page_roles";
CREATE TABLE IF NOT EXISTS "page_roles" (
	"page_id"	integer,
	"role_id"	integer,
	PRIMARY KEY("page_id","role_id"),
	CONSTRAINT "fk_page_roles_role" FOREIGN KEY("role_id") REFERENCES "roles"("id") ON DELETE SET NULL ON UPDATE CASCADE,
	CONSTRAINT "fk_page_roles_page" FOREIGN KEY("page_id") REFERENCES "pages"("id") ON DELETE SET NULL ON UPDATE CASCADE
);
DROP TABLE IF EXISTS "user_roles constraint:_on_update:_cascad_e,_on_delete:_cascades";
CREATE TABLE IF NOT EXISTS "user_roles constraint:_on_update:_cascad_e,_on_delete:_cascades" (
	"role_id"	integer,
	"user_id"	integer,
	PRIMARY KEY("role_id","user_id"),
	CONSTRAINT "fk_user_roles constraint:_on_update:_cascad_e,_on_delete969b5d96" FOREIGN KEY("user_id") REFERENCES "users"("id"),
	CONSTRAINT "fk_user_roles constraint:_on_update:_cascad_e,_on_delete09e01c66" FOREIGN KEY("role_id") REFERENCES "roles"("id")
);
DROP TABLE IF EXISTS "user_roles constraint:_on_update:_cascades";
CREATE TABLE IF NOT EXISTS "user_roles constraint:_on_update:_cascades" (
	"role_id"	integer,
	"user_id"	integer,
	CONSTRAINT "fk_user_roles constraint:_on_update:_cascades_role" FOREIGN KEY("role_id") REFERENCES "roles"("id"),
	CONSTRAINT "fk_user_roles constraint:_on_update:_cascades_user" FOREIGN KEY("user_id") REFERENCES "users"("id"),
	PRIMARY KEY("role_id","user_id")
);
DROP TABLE IF EXISTS "end_points";
CREATE TABLE IF NOT EXISTS "end_points" (
	"id"	integer,
	"name"	text UNIQUE,
	"route_paths"	text,
	"description"	text,
	"feature_id"	integer,
	"method"	text,
	CONSTRAINT "fk_features_endpoints" FOREIGN KEY("feature_id") REFERENCES "features"("id"),
	PRIMARY KEY("id")
);
INSERT INTO "roles" ("id","name","description","active","app_id") VALUES (1,'superuser','Have Access to All resources in the Apps',1,2),
 (2,'standard','Have Access to Limited Features',1,1),
 (3,'administrator','Have Access System Admin Privileges',1,1),
 (4,'app_role','Have App CURD abilities ',1,1),
 (5,'new role','add new role',1,NULL),
 (6,'work','add new new working for all',1,NULL),
 (7,'testing WAl','Checking WAL again',1,NULL);
INSERT INTO "user_roles" ("user_id","role_id") VALUES (1,1),
 (1,2),
 (1,3),
 (2,1),
 (2,2),
 (2,3),
 (5,1),
 (5,3),
 (6,2),
 (6,1);
INSERT INTO "pages" ("id","name","active","description") VALUES (1,'Login',1,'Login'),
 (2,'Home',1,'Home'),
 (3,'View Roles',1,'View Roles'),
 (4,'View Users',1,'View Users'),
 (5,'View Pages',1,'View Pages'),
 (6,'View Features',1,'View Features'),
 (7,'View Endpoints',1,'View Endpoints'),
 (8,'View Apps',1,'View Apps'),
 (9,'Test',1,'Testing Page'),
 (10,'Test UI',0,'Adding From UI'),
 (11,'Another',0,'From UI intact Fine');
INSERT INTO "features" ("id","name","description","active","role_id") VALUES (1,'login','login',1,1),
 (2,'check_login','check_login',1,1),
 (3,'create_user','create_user',1,NULL),
 (4,'get_users_list','get_users_list',1,2),
 (5,'update_user_detail','update_user_detail',1,2),
 (6,'delete_user','delete_user',1,3),
 (7,'activate_deactivate_user','activate_deactivate_user',1,2),
 (8,'get_user_roles','get_user_roles',1,3),
 (9,'update_user_role','update_user_role',1,3),
 (10,'delete_user_role','delete_user_role',1,NULL),
 (11,'create_role','create_role',1,1),
 (12,'update_role','update_role',1,2),
 (13,'delete_role','delete_role',1,3),
 (14,'activate_deactivate_role','activate_deactivate_role',1,2),
 (15,'create_feature','create_feature',1,NULL),
 (16,'update_feature_details','update_feature_details',1,3),
 (17,'activate_deactivate_feature','activate_deactivate_feature',1,2),
 (18,'delete_feature','delete_feature',1,3),
 (19,'map_feature_with_role','map_feature_with_role',1,NULL),
 (20,'map_feature_with_endpoints','map_feature_with_endpoints',1,NULL),
 (21,'auto_populate','auto_populate',1,3),
 (22,'get_list_of_endpoints','get_list_of_endpoints',1,3),
 (23,'create_page','create_page',1,3),
 (24,'activate_deactivate_page','activate_deactivate_page',1,3),
 (25,'map_features_with_page','map_features_with_page',1,3),
 (26,'create_app','create_app',1,3),
 (27,'activate_deactivate_app','activate_deactivate_app',1,3),
 (28,'map_features_with_app','map_features_with_app',1,3),
 (29,'get_app_secrete','get_app_secrete',1,3),
 (30,'list_app_users','list_app_users',1,4),
 (31,'Some Feature','Working for tests one two three',1,NULL),
 (32,'new feature','Testing feature Edited Again',0,NULL),
 (33,'somenew','checking it works with mobile',1,NULL);
INSERT INTO "apps" ("id","name","uuid","active","description") VALUES (1,'BlueAdmin','ddae04a8-7efb-4c80-bd34-eb2fb41449b7',1,'SSO Role Based User Administration solution'),
 (2,'Super App','5cadd46f-d823-463f-8b43-291f9b8eb49b',1,'For Testing Developer Mode Access Edited'),
 (3,'App UI','55393178-008e-4c17-bb89-1e2b7a2a703d',1,'Add from UI'),
 (4,'App UI 2','e9b7d7e2-8dda-4762-ab85-8d33a4efb036',1,'Add from UI'),
 (5,'Add UI','9b548078-b484-49d6-ab40-6624f16f09a9',1,'second attempt'),
 (6,'Mobile  ','70dacd97-6327-4e52-b178-3f7826bfd1e5',1,'For Mob Section Updated'),
 (7,'Logging','151b0db4-b80d-48f4-8164-093524525078',1,'Check Logging From UI'),
 (8,'Inactive Role','bee4290e-d58c-450d-8761-fbd5008d8196',0,'Check Inactive');
INSERT INTO "users" ("id","uuid","email","password","date_registered","disabled") VALUES (1,'7e73750b-08db-4281-8e34-bc51837d70c6','beimdegefu@gmail.com','089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a','2023-12-14 11:35:31',1),
 (2,'112fe86f-5064-4093-ba6d-797850d04e28','bushubeke@gmail.com','089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a','2023-12-14 12:43:07',0),
 (3,'2e16400d-d8a7-415b-b6bf-d6129dfb6ef7','bushudegefu@gmail.com','089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a','2023-12-14 12:47:08',0),
 (4,'18925651-a0ae-4028-adf3-fe4392f703a9','bushudegefu2@gmail.com','089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a','2023-12-14 12:47:45',0),
 (5,'622c0ccf-7d23-43d7-b559-aa07cd136076','bushudegefu3@gmail.com','089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a','2023-12-14 13:42:30',1),
 (6,'5e7f9e7e-c1d7-4e36-9dd9-a3c796f296a3','eyobamare@yahoo.com','089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a','2023-12-18 06:15:12',1);
INSERT INTO "page_roles" ("page_id","role_id") VALUES (1,2),
 (1,1);
INSERT INTO "end_points" ("id","name","route_paths","description","feature_id","method") VALUES (1,'swagger_routes_get','/docs/*','swagger_routes-GET',NULL,'GET'),
 (2,'custom_metrics_route_get','/lmetrics','custom_metrics_route-GET',9,'GET'),
 (3,'check_login_get','/api/v1/checklogin','check_login-GET',NULL,'GET'),
 (4,'roles_get','/api/v1/roles','roles-GET',NULL,'GET'),
 (5,'roles_single_get','/api/v1/roles/:id','roles_single-GET',NULL,'GET'),
 (6,'drop_roles_get','/api/v1/droproles','drop_roles-GET',NULL,'GET'),
 (7,'roles_endpoints_get','/api/v1/role_endpoints','roles_endpoints-GET',NULL,'GET'),
 (8,'features_get','/api/v1/features','features-GET',NULL,'GET'),
 (9,'features_single_get','/api/v1/features/:id','features_single-GET',NULL,'GET'),
 (10,'drop_features_get','/api/v1/featuredrop','drop_features-GET',8,'GET'),
 (11,'apps_get','/api/v1/apps','apps-GET',8,'GET'),
 (12,'apps_single_get','/api/v1/apps/:id','apps_single-GET',9,'GET'),
 (13,'users_get','/api/v1/users','users-GET',4,'GET'),
 (14,'user_single_get','/api/v1/users/:id','user_single-GET',4,'GET'),
 (15,'get_user_roles_get','/api/v1/userrole/:user_id','get_user_roles-GET',NULL,'GET'),
 (16,'pages_get','/api/v1/pages','pages-GET',NULL,'GET'),
 (17,'page_single_get','/api/v1/pages/:id','page_single-GET',NULL,'GET'),
 (18,'get_page_roles_get','/api/v1/pagesroles/:page_id','get_page_roles-GET',NULL,'GET'),
 (19,'end_point_get','/api/v1/endpoints','end_point-GET',9,'GET'),
 (20,'end_point_single_get','/api/v1/endpoints/:id','end_point_single-GET',NULL,'GET'),
 (21,'drop_endpoints_get','/api/v1/endpointdrop','drop_endpoints-GET',8,'GET'),
 (22,'login_route_post','/api/v1/login','login_route-POST',NULL,'POST'),
 (23,'roles_post','/api/v1/roles','roles-POST',NULL,'POST'),
 (24,'features_post','/api/v1/features','features-POST',NULL,'POST'),
 (25,'apps_post','/api/v1/apps','apps-POST',8,'POST'),
 (26,'users_post','/api/v1/users','users-POST',NULL,'POST'),
 (27,'user_role_post','/api/v1/userrole/:user_id/:role_id','user_role-POST',4,'POST'),
 (28,'pages_post','/api/v1/pages','pages-POST',NULL,'POST'),
 (29,'page_roles_post','/api/v1/pagerole/:page_id/:role_id','page_roles-POST',NULL,'POST'),
 (30,'end_point_post','/api/v1/endpoints','end_point-POST',NULL,'POST'),
 (31,'send_email_post','/api/v1/email','send_email-POST',NULL,'POST'),
 (32,'activate_deactivate_role_put','/api/v1/roles/:role_id','activate_deactivate_role-PUT',NULL,'PUT'),
 (33,'activate_deactivate_features_put','/api/v1/features/:feature_id','activate_deactivate_features-PUT',NULL,'PUT'),
 (34,'activate_deactivate_user_put','/api/v1/users/:user_id','activate_deactivate_user-PUT',NULL,'PUT'),
 (35,'roles_single_delete','/api/v1/roles/:id','roles_single-DELETE',NULL,'DELETE'),
 (36,'features_single_delete','/api/v1/features/:id','features_single-DELETE',NULL,'DELETE'),
 (37,'feature_role_delete','/api/v1/featuresrole/:feature_id','feature_role-DELETE',NULL,'DELETE'),
 (38,'apps_single_delete','/api/v1/apps/:id','apps_single-DELETE',32,'DELETE'),
 (39,'user_single_delete','/api/v1/users/:id','user_single-DELETE',NULL,'DELETE'),
 (40,'user_role_delete','/api/v1/userrole/:user_id/:role_id','user_role-DELETE',NULL,'DELETE'),
 (41,'page_single_delete','/api/v1/pages/:id','page_single-DELETE',NULL,'DELETE'),
 (42,'page_roles_delete','/api/v1/pagerole/:page_id/:role_id','page_roles-DELETE',NULL,'DELETE'),
 (43,'end_point_single_delete','/api/v1/endpoints/:id','end_point_single-DELETE',NULL,'DELETE'),
 (44,'feature_endpoint_delete','/api/v1/feature_endpoint/:endpoint_id','feature_endpoint-DELETE',NULL,'DELETE'),
 (45,'roles_single_patch','/api/v1/roles/:id','roles_single-PATCH',NULL,'PATCH'),
 (46,'roles_app_patch','/api/v1/approle/{role_id}','roles_app-PATCH',NULL,'PATCH'),
 (47,'features_single_patch','/api/v1/features/:id','features_single-PATCH',NULL,'PATCH'),
 (48,'feature_role_patch','/api/v1/featuresrole/:feature_id','feature_role-PATCH',NULL,'PATCH'),
 (49,'apps_single_patch','/api/v1/apps/:id','apps_single-PATCH',9,'PATCH'),
 (50,'user_single_patch','/api/v1/users/:id','user_single-PATCH',4,'PATCH'),
 (51,'page_single_patch','/api/v1/pages/:id','page_single-PATCH',NULL,'PATCH'),
 (52,'end_point_single_patch','/api/v1/endpoints/:id','end_point_single-PATCH',NULL,'PATCH'),
 (53,'feature_endpoint_patch','/api/v1/feature_endpoint/:endpoint_id','feature_endpoint-PATCH',NULL,'PATCH'),
 (54,'apps_delete','/api/v1/featuresrole/:feature_id','apps-DELETE',NULL,'DELETE'),
 (55,'apps_patch','/api/v1/featuresrole/:feature_id','apps-PATCH',NULL,'PATCH'),
 (56,'apps_features_get','/api/v1/appsmatrix/:id','apps_features-GET',NULL,'GET'),
 (57,'for test','/somepathchanged','Checking update',NULL,'');
COMMIT;
