
INSERT INTO apps (id, name, uuid, active, description) VALUES 
  	(1,'BlueAdmin','48015a9b-5a86-4a15-944b-94108aa78b4b',true,'SSO Role Based User Administration solution'),
  	(2,'BlueCom','21028fa1-8e04-464e-8be7-95ea9c82994b',true,'Commercial Order Management'),
	(3, 'BlueHRM','9359b1ba-98b5-427c-96d8-023fc33cd1b0',true,'Human Resource Management') ON CONFLICT DO NOTHING;


INSERT INTO pages (id, name, active, description) VALUES
	(1, 'Login', true, 'Login'),
	(2, 'Home', true, 'Home'),
	(3, 'Role', true, 'View Roles'),
	(4, 'User', true, 'View Users'),
	(5, 'Page', true, 'View Pages'),
	(6, 'Feature', true, 'View Features'),
	(7, 'Endpoint', true, 'View Endpoints'),
	(8, 'App', true, 'View Apps'),
	(9, 'Sign Up', true, 'Sign Up Page For users Who have not registered') ON CONFLICT DO NOTHING;

INSERT INTO roles (id, name, description, active, app_id) VALUES
	(1, 'superuser', 'Have Access to All resources in the Apps', true, 1),
	(2, 'standard', 'Have Access to Limited Features', true, 1),
	(3, 'administrator', 'Have Access System Admin Privileges', true, 1),
	(4, 'app_role', 'Have App CURD abilities', true, 1),
	(5, 'Anonymous', 'For Pages that do not need user sign in', true, 1),
	(6, 'Drop Down ', 'To access Drop Down menu fetching  endpoints', true, 1) ON CONFLICT DO NOTHING;

INSERT INTO users (id, uuid, email, password, date_registered, disabled) VALUES
	(4, '8a200cd4-9067-4508-b93c-4b242ef03740', 'beimnet.degefu@gmail.com', '089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a', '2024-03-08 16:49:53.754679+00', false),
	(6, '79731de0-de10-4f62-bcc5-7637d132110a', 'mickyasne123@gmail.com', '3d5267ceaa0759b8756837e4e817517fae8fa2d168267053e78428ac0941d5542c7bee61536be21f7d117fed80bc9aa27e2d450e261ab361ece2a701a1e2da72', '2024-03-12 10:52:01.052905+00', false),
	(7, '586c2bc7-9a66-49df-a81e-93b2efaec8c9', 'mickyasne12@gmail.com', 'adfebcc24b8ce4be96b83381f08275bbd6fe355dee6a1a1407247cd15bbb8ec5952b5680af51c532d8785f950e4bb7ddeab72efc737c4656bcd27b7f73b72bf1', '2024-03-12 12:03:48.01416+00', false),
	(1, '38ca7360-0138-4b0f-8985-b307ad188e92', 'superuser@mail.com', '089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a', '2024-01-16 08:27:55.628203+00', false),
	(2, '1323b2d9-5755-4e4c-9af5-93c17f59e6fd', 'standarduser@mail.com', '089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a', '2024-01-17 12:33:56.443849+00', false),
	(3, '12bc4954-487b-4263-8bdc-cacbf720f623', 'adminuser@mail.com', '089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a', '2024-03-01 06:58:25.299887+00', true),
	(5, 'b5c4d708-71de-4384-a91d-73843bb45947', 'somesuper@gmail.com', 'c320a2866e3a6feb0a6abf607c55f0466873c34a68c27fb6dd5e16fc53173f96f2326b97ec943063071fb95f14f883701a7b827e8b18767754190b623e051111', '2024-03-12 10:48:50.626919+00', false) ON CONFLICT DO NOTHING;



INSERT INTO features (id, name, description, active, role_id) VALUES
	(1, 'role_read', 'View List of Roles', true, 2),
	(2, 'role_write', 'Privilege to Update ,Create, View and Delete  ', true, 3),
	(3, 'user_read', 'View List of Users', true, 2),
	(4, 'user_write', 'Privilege to Update ,Create, View and Delete  ', true, 3),
	(5, 'page_read', 'View List of Pages', true, 2),
	(6, 'page_write', 'Privilege to Update ,Create, View and Delete  ', true, 3),
	(7, 'app_read', 'View List of Apps', true, 2),
	(8, 'app_write', 'Privilege to Update ,Create, View and Delete ', true, 4),
	(9, 'endpoint_read', 'View List of Endpoints', true, 2),
	(10, 'endpoint_write', 'Privilege to Update ,Create, View and Delete  ', true, 3),
	(11, 'feature_read', 'View List of Features', true, 2),
	(12, 'features_write', 'Privilege to Update ,Create, View and Delete  ', true, 3),
	(13, 'login', 'Endpoints that can be accessed with out Logging In', true, 5),
	(14, 'drop_down', 'Endpoints that fetch Drowns ', false, 6) ON CONFLICT DO NOTHING;


INSERT INTO end_points (name, route_paths, method, description, feature_id) VALUES
	('swagger_routes_get', '/docs/*', 'GET', 'swagger_routes-GET', NULL),
	('custom_metrics_route_get', '/lmetrics', 'GET', 'custom_metrics_route-GET', 13),
	('check_login_get', '/api/v1/checklogin', 'GET', 'check_login-GET', 13),
	('roles_get', '/api/v1/roles', 'GET', 'roles-GET', 1),
	('roles_single_get', '/api/v1/roles/:id', 'GET', 'roles_single-GET', 1),
	('drop_roles_get', '/api/v1/droproles', 'GET', 'drop_roles-GET', 14),
	('roles_endpoints_get', '/api/v1/role_endpoints', 'GET', 'roles_endpoints-GET', 1),
	('features_get', '/api/v1/features', 'GET', 'features-GET', 11),
	('features_single_get', '/api/v1/features/:id', 'GET', 'features_single-GET', 11),
	('drop_features_get', '/api/v1/featuredrop', 'GET', 'drop_features-GET', 14),
	('apps_get', '/api/v1/apps', 'GET', 'apps-GET', 7),
	('apps_single_get', '/api/v1/apps/:id', 'GET', 'apps_single-GET', 7),
	('drop_sppd_get', '/api/v1/appsdrop', 'GET', 'drop_sppd-GET', 14),
	('apps_features_get', '/api/v1/appsmatrix/:id', 'GET', 'apps_features-GET', 7),
	('users_get', '/api/v1/users', 'GET', 'users-GET', 3),
	('user_single_get', '/api/v1/users/:id', 'GET', 'user_single-GET', 3),
	('get_user_roles_get', '/api/v1/userrole/:user_id', 'GET', 'get_user_roles-GET', 3),
	('pages_get', '/api/v1/pages', 'GET', 'pages-GET', 5),
	('page_single_get', '/api/v1/pages/:id', 'GET', 'page_single-GET', 5),
	('get_page_roles_get', '/api/v1/pagesroles/:page_id', 'GET', 'get_page_roles-GET', 5),
	('end_point_get', '/api/v1/endpoints', 'GET', 'end_point-GET', 9),
	('end_point_single_get', '/api/v1/endpoints/:id', 'GET', 'end_point_single-GET', 9),
	('drop_endpoints_get', '/api/v1/endpointdrop', 'GET', 'drop_endpoints-GET', 14),
	('dashboard_get', '/api/v1/dashboard', 'GET', 'dashboard-GET', 14),
	('login_route_post', '/api/v1/login', 'POST', 'login_route-POST', 13),
	('roles_post', '/api/v1/roles', 'POST', 'roles-POST', 2),
	('features_post', '/api/v1/features', 'POST', 'features-POST', 12),
	('apps_post', '/api/v1/apps', 'POST', 'apps-POST', 8),
	('users_post', '/api/v1/users', 'POST', 'users-POST', 4),
	('user_role_post', '/api/v1/userrole/:user_id/:role_id', 'POST', 'user_role-POST', 4),
	('pages_post', '/api/v1/pages', 'POST', 'pages-POST', 6),
	('page_roles_post', '/api/v1/pagerole/:page_id/:role_id', 'POST', 'page_roles-POST', 6),
	('end_point_post', '/api/v1/endpoints', 'POST', 'end_point-POST', 10),
	('send_email_post', '/api/v1/email', 'POST', 'send_email-POST', NULL),
	('blob_picture_post', '/api/v1/blobpic', 'POST', 'blob_picture-POST', 13),
	('blob_video_post', '/api/v1/blobvideo', 'POST', 'blob_video-POST', 13),
	('activate_deactivate_role_put', '/api/v1/roles/:role_id', 'PUT', 'activate_deactivate_role-PUT', 2),
	('activate_deactivate_features_put', '/api/v1/features/:feature_id', 'PUT', 'activate_deactivate_features-PUT', 12),
	('activate_deactivate_user_put', '/api/v1/users/:user_id', 'PUT', 'activate_deactivate_user-PUT', 4),
	('roles_single_delete', '/api/v1/roles/:id', 'DELETE', 'roles_single-DELETE', 2),
	('features_single_delete', '/api/v1/features/:id', 'DELETE', 'features_single-DELETE', 12),
	('feature_role_delete', '/api/v1/featuresrole/:feature_id', 'DELETE', 'feature_role-DELETE', 12),
	('apps_single_delete', '/api/v1/apps/:id', 'DELETE', 'apps_single-DELETE', 8),
	('user_single_delete', '/api/v1/users/:id', 'DELETE', 'user_single-DELETE', 4),
	('user_role_delete', '/api/v1/userrole/:user_id/:role_id', 'DELETE', 'user_role-DELETE', 4),
	('page_single_delete', '/api/v1/pages/:id', 'DELETE', 'page_single-DELETE', 6),
	('page_roles_delete', '/api/v1/pagerole/:page_id/:role_id', 'DELETE', 'page_roles-DELETE', 6),
	('end_point_single_delete', '/api/v1/endpoints/:id', 'DELETE', 'end_point_single-DELETE', 10),
	('feature_endpoint_delete', '/api/v1/feature_endpoint/:endpoint_id', 'DELETE', 'feature_endpoint-DELETE', 12),
	('roles_single_patch', '/api/v1/roles/:id', 'PATCH', 'roles_single-PATCH', 2),
	('roles_app_patch', '/api/v1/approle/:role_id', 'PATCH', 'roles_app-PATCH', 2),
	('features_single_patch', '/api/v1/features/:id', 'PATCH', 'features_single-PATCH', 12),
	('feature_role_patch', '/api/v1/featuresrole/:feature_id', 'PATCH', 'feature_role-PATCH', 12),
	('apps_single_patch', '/api/v1/apps/:id', 'PATCH', 'apps_single-PATCH', 8),
	('user_single_patch', '/api/v1/users/:id', 'PATCH', 'user_single-PATCH', 4),
	('page_single_patch', '/api/v1/pages/:id', 'PATCH', 'page_single-PATCH', 6),
	('end_point_single_patch', '/api/v1/endpoints/:id', 'PATCH', 'end_point_single-PATCH', 10),
	('feature_endpoint_patch', '/api/v1/feature_endpoint/:endpoint_id', 'PATCH', 'feature_endpoint-PATCH', 12),
	('change_reset_password_put', '/api/v1/users/:email_id', 'PUT', 'change_reset_password-PUT', 4) ON CONFLICT DO NOTHING;


INSERT INTO page_roles (page_id, role_id) VALUES
	(1, 5),
	(5, 3),
	(2, 5),
	(3, 3),
	(3, 1),
	(2, 1),
	(1, 1),
	(4, 2),
	(4, 3),
	(4, 1),
	(5, 2),
	(5, 1),
	(6, 1),
	(6, 2),
	(6, 3),
	(7, 1),
	(7, 3),
	(8, 1),
	(8, 3),
	(7, 2),
	(3, 2),
	(8, 2),
	(2, 2),
	(9, 1) ON CONFLICT DO NOTHING;


INSERT INTO user_roles (role_id, user_id) VALUES
	(1, 1),
	(1, 4),
	(1, 5),
	(2, 2),
	(2, 6),
	(2, 7),
	(3, 3) ON CONFLICT DO NOTHING;

INSERT INTO jwt_salts (id, salt_a, salt_b) VALUES
	(1, 'ANJHTrDA7guiAaE', 'wnQh26QQNm9Oc0x') ON CONFLICT DO NOTHING;
