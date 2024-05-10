INSERT INTO rate_limit_rule (rl_rule_id, max_shipments, time_shipment)
VALUES (1, 5, 'W');

INSERT INTO rate_limit_rule (rl_rule_id, max_shipments, time_shipment)
VALUES (2, 3, 'M');

INSERT INTO rate_limit_rule (rl_rule_id, max_shipments, time_shipment)
VALUES (3, 12, 'Y');

INSERT INTO rate_limit_rule (rl_rule_id, max_shipments, time_shipment)
VALUES (4, 2, 'h');

INSERT INTO rate_limit_rule (rl_rule_id, max_shipments, time_shipment)
VALUES (5, 1, 'm');

INSERT INTO type_notification (type_id, rl_rule_id)
VALUES ('status', 2);

INSERT INTO type_notification (type_id, rl_rule_id)
VALUES ('news', 4);

INSERT INTO type_notification (type_id, rl_rule_id)
VALUES ('marketing', 3);