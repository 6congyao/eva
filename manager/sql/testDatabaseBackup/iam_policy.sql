INSERT INTO public.iam_policy (policy_id, policy_name, description, qrn, path, statement, version, create_time, root_user_id) VALUES ('iamp-cy6bTxYp', 'K8SAccess', 'This policy was created for k8s', 'qrn:partition::iam:usr-Vtl3VCfF:policy/K8SAccess', '/', '{
		"Version": "2018-5-5",
		"statement": [
			{
				"effect": "allow",
				"action": [
					"k8s:list",
					"k8s:get"
				],
				"resource": "k8s:pods"
			},
			{
				"effect": "allow",
				"action": "k8s:watch",
				"resource": [
					"k8s:pods",
					"k8s:pods/log"
				]
			}
		]
	}', '2018-05-04', '2018-05-16 10:04:50.179063', 'usr-Vtl3VCfF');
INSERT INTO public.iam_policy (policy_id, policy_name, description, qrn, path, statement, version, create_time, root_user_id) VALUES ('iamp-be0rrgur', 'k8sadmin', 'kubenetes', 'qrn:partition::iam:admin:policy/k8sadmin', '/', '{"Version":"2012-04-01","Statement":[{"action":"k8s:list","resource":"k8s:pods","effect":"allow"}]}', '2012-04-01', '2018-05-17 16:56:39.697037', 'admin');
INSERT INTO public.iam_policy (policy_id, policy_name, description, qrn, path, statement, version, create_time, root_user_id) VALUES ('iamp-tft72s3p', 'QstorAccess', null, 'qrn:partition::iam:admin:policy/QstorAccess', '/', '{"Version":"2012-04-01","Statement":[{"action":["qstor:modify","qstor:delete"],"resource":"qstor:*","effect":"deny"}]}', '2012-04-01', '2018-05-17 16:57:09.626968', 'admin');
INSERT INTO public.iam_policy (policy_id, policy_name, description, qrn, path, statement, version, create_time, root_user_id) VALUES ('iamp-cy6bTxYn', 'IAMPolicyAccess', 'This policy was created for iam service', 'qrn:partition::iam:usr-Vtl3VCfF:policy/IAMPolicyAccess', '/', '{
		"Version": "2018-5-4",
		"statement": [
			{
				"effect": "allow",
				"action": [
					"iam:CreatePolicy",
					"iam:DeletePolicy"
				],
				"resource": "*"
			},
			{
				"effect": "allow",
				"action": "iam:CreateRole",
				"resource": "*"
			}
		]
	}', '2018-05-04', '2018-05-16 10:04:50.179063', 'usr-Vtl3VCfF');