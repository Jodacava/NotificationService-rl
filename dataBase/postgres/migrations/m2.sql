insert into rate_limit_rule (rl_rule_id, max_shipments, time_shipment, type_id)
values  (3, 12, 'Y', 'invitation'),
        (4, 2, 'h', 'news'),
        (2, 3, 'M', 'marketing'),
        (1, 5, 'W', 'status'),
        (5, 1, 'm', 'position');

insert into type_notification (type_id, rl_rule_id)
values  ('status', 2),
        ('news', 4),
        ('marketing', 3);