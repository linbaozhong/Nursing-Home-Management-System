CREATE TABLE `assessment` (
                              `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                              `tenant_id` bigint NOT NULL COMMENT '租户ID',
                              `elderly_id` bigint DEFAULT NULL COMMENT '老人ID',
                              `assessment_date` date NOT NULL COMMENT '评估日期',
                              `assessment_type` tinyint NOT NULL COMMENT '评估类型：1-能力评估，2-健康评估',
                              `scale_data` json DEFAULT NULL COMMENT '量表原始数据',
                              `result_level` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '评估结果（护理等级）',
                              `evaluator_id` bigint DEFAULT NULL COMMENT '评估人ID',
                              `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                              `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                              PRIMARY KEY (`id`),
                              KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  KEY `idx_evaluator_id` (`evaluator_id`),
  CONSTRAINT `fk_assessments_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_assessments_staff` FOREIGN KEY (`evaluator_id`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_assessments_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评估记录表';

CREATE TABLE `bed` (
                       `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                       `tenant_id` bigint NOT NULL COMMENT '租户ID',
                       `room_id` bigint NOT NULL COMMENT '房间ID',
                       `bed_no` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '床位编号',
                       `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-空置，1-占用，2-维修',
                       `type` tinyint NOT NULL DEFAULT '1' COMMENT '类型：1-普通床，2-护理床，3-加床',
                       `price` bigint DEFAULT '0' COMMENT '床位费（分）',
                       `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                       `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                       `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                       PRIMARY KEY (`id`),
                       KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_room_id` (`room_id`),
  CONSTRAINT `fk_beds_room` FOREIGN KEY (`room_id`) REFERENCES `room` (`id`),
  CONSTRAINT `fk_beds_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='床位表';

CREATE TABLE `bill` (
                        `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                        `tenant_id` bigint NOT NULL COMMENT '租户ID',
                        `elderly_id` bigint NOT NULL COMMENT '老人ID',
                        `bill_no` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账单编号',
                        `bill_period` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账单周期',
                        `total_amount` bigint NOT NULL DEFAULT '0' COMMENT '总金额（分）',
                        `paid_amount` bigint NOT NULL DEFAULT '0' COMMENT '已付金额（分）',
                        `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-未支付，1-部分支付，2-已支付，3-逾期',
                        `due_date` date DEFAULT NULL COMMENT '缴费截止日',
                        `generated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生成时间',
                        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `uk_bill_no` (`bill_no`),
                        KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  KEY `idx_status` (`status`),
  KEY `idx_bill_period` (`bill_period`),
  CONSTRAINT `fk_bills_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_bills_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='账单表';

CREATE TABLE `bill_item` (
                             `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                             `bill_id` bigint NOT NULL COMMENT '账单ID',
                             `fee_item_id` bigint DEFAULT NULL COMMENT '费用项ID',
                             `description` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '描述',
                             `quantity` decimal(10,2) NOT NULL COMMENT '数量',
                             `unit_price` bigint NOT NULL DEFAULT '0' COMMENT '单价（分）',
                             `amount` bigint NOT NULL DEFAULT '0' COMMENT '小计（分）',
                             `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                             PRIMARY KEY (`id`),
                             KEY `idx_bill_id` (`bill_id`),
  KEY `idx_fee_item_id` (`fee_item_id`),
  CONSTRAINT `fk_bill_items_bill` FOREIGN KEY (`bill_id`) REFERENCES `bill` (`id`),
  CONSTRAINT `fk_bill_items_fee_item` FOREIGN KEY (`fee_item_id`) REFERENCES `fee_item` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='账单明细表';

CREATE TABLE `care_item` (
                             `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                             `tenant_id` bigint NOT NULL COMMENT '租户ID',
                             `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '项目名称',
                             `category` tinyint NOT NULL COMMENT '类别：1-基础护理，2-医疗护理，3-康复',
                             `default_time` time DEFAULT NULL COMMENT '默认执行时间',
                             `need_photo` tinyint NOT NULL DEFAULT '0' COMMENT '是否需要拍照确认：0-否，1-是',
                             `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                             `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                             `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                             PRIMARY KEY (`id`),
                             KEY `idx_tenant_id` (`tenant_id`),
  CONSTRAINT `fk_care_items_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='护理项目表';

CREATE TABLE `care_level` (
                              `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                              `tenant_id` bigint NOT NULL COMMENT '租户ID',
                              `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '等级名称',
                              `daily_price` bigint NOT NULL DEFAULT '0' COMMENT '每日护理费（分）',
                              `service_items` json DEFAULT NULL COMMENT '默认服务项目ID列表',
                              `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                              `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                              PRIMARY KEY (`id`),
                              KEY `idx_tenant_id` (`tenant_id`),
  CONSTRAINT `fk_care_levels_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='护理等级表';

CREATE TABLE `care_plan` (
                             `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                             `tenant_id` bigint NOT NULL COMMENT '租户ID',
                             `elderly_id` bigint NOT NULL COMMENT '老人ID',
                             `care_item_id` bigint NOT NULL COMMENT '护理项目ID',
                             `exec_time` time DEFAULT NULL COMMENT '执行时间',
                             `frequency` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '执行频率（每天/每周几/隔天）',
                             `start_date` date NOT NULL COMMENT '生效开始日',
                             `end_date` date DEFAULT NULL COMMENT '生效结束日',
                             `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：1-生效，0-暂停',
                             `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                             `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                             `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                             PRIMARY KEY (`id`),
                             KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  KEY `idx_care_item_id` (`care_item_id`),
  CONSTRAINT `fk_care_plans_care_item` FOREIGN KEY (`care_item_id`) REFERENCES `care_item` (`id`),
  CONSTRAINT `fk_care_plans_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_care_plans_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='护理计划表';

CREATE TABLE `care_task` (
                             `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                             `tenant_id` bigint NOT NULL COMMENT '租户ID',
                             `elderly_id` bigint NOT NULL COMMENT '老人ID',
                             `care_item_id` bigint NOT NULL COMMENT '护理项目ID',
                             `plan_id` bigint DEFAULT NULL COMMENT '关联的护理计划ID',
                             `scheduled_time` datetime NOT NULL COMMENT '计划执行时间',
                             `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-待执行，1-已完成，2-超时未完成，3-已跳过',
                             `actual_time` datetime DEFAULT NULL COMMENT '实际完成时间',
                             `executor_id` bigint DEFAULT NULL COMMENT '执行人ID',
                             `photo_url` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '照片地址',
                             `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                             `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                             `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                             PRIMARY KEY (`id`),
                             KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  KEY `idx_plan_id` (`plan_id`),
  KEY `idx_executor_id` (`executor_id`),
  KEY `idx_scheduled_time` (`scheduled_time`),
  KEY `idx_status` (`status`),
  KEY `fk_care_tasks_care_item` (`care_item_id`),
  CONSTRAINT `fk_care_tasks_care_item` FOREIGN KEY (`care_item_id`) REFERENCES `care_item` (`id`),
  CONSTRAINT `fk_care_tasks_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_care_tasks_executor` FOREIGN KEY (`executor_id`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_care_tasks_plan` FOREIGN KEY (`plan_id`) REFERENCES `care_plan` (`id`),
  CONSTRAINT `fk_care_tasks_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='护理任务表';

CREATE TABLE `checkout` (
                            `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                            `tenant_id` bigint NOT NULL COMMENT '租户ID',
                            `elderly_id` bigint NOT NULL COMMENT '老人ID',
                            `checkout_date` date NOT NULL COMMENT '退住日期',
                            `reason` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '原因',
                            `settlement_amount` bigint NOT NULL DEFAULT '0' COMMENT '结算金额（分，正退费，负补缴）',
                            `operator_id` bigint DEFAULT NULL COMMENT '操作人ID',
                            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            PRIMARY KEY (`id`),
                            KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  CONSTRAINT `fk_checkouts_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_checkouts_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='退住记录表';

CREATE TABLE `consultation` (
                                `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                `elderly_name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '老人姓名',
                                `contact_phone` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '联系电话',
                                `channel` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '来源渠道',
                                `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-跟进中，1-已入住，2-放弃',
                                `follow_up` text COLLATE utf8mb4_unicode_ci COMMENT '跟进记录',
                                `intended_move_in_date` date DEFAULT NULL COMMENT '意向入住日期',
                                `created_by` bigint DEFAULT NULL COMMENT '创建人ID',
                                `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                                PRIMARY KEY (`id`),
                                KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_created_by` (`created_by`),
  CONSTRAINT `fk_consultations_staff` FOREIGN KEY (`created_by`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_consultations_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='咨询记录表';

CREATE TABLE `contract` (
                            `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                            `tenant_id` bigint NOT NULL COMMENT '租户ID',
                            `elderly_id` bigint NOT NULL COMMENT '老人ID',
                            `contract_no` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '合同编号',
                            `file_url` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '电子合同文件地址',
                            `start_date` date NOT NULL COMMENT '生效日期',
                            `end_date` date DEFAULT NULL COMMENT '结束日期',
                            `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：1-有效，2-到期，3-终止',
                            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `uk_contract_no` (`contract_no`),
                            KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  CONSTRAINT `fk_contracts_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_contracts_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='合同表';

CREATE TABLE `deposit` (
                           `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                           `tenant_id` bigint NOT NULL COMMENT '租户ID',
                           `elderly_id` bigint NOT NULL COMMENT '老人ID',
                           `amount` bigint NOT NULL DEFAULT '0' COMMENT '押金金额（分）',
                           `balance` bigint NOT NULL DEFAULT '0' COMMENT '当前余额（分）',
                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '收取时间',
                           `refunded_at` datetime DEFAULT NULL COMMENT '退还时间',
                           `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           PRIMARY KEY (`id`),
                           KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  CONSTRAINT `fk_deposits_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_deposits_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='押金记录表';

CREATE TABLE `drug_inventory` (
                                  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                  `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                  `drug_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '药品名称',
                                  `batch_no` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '批号',
                                  `expiry_date` date DEFAULT NULL COMMENT '有效期',
                                  `quantity` int NOT NULL DEFAULT '0' COMMENT '当前库存',
                                  `unit` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '单位',
                                  `price` bigint DEFAULT '0' COMMENT '进价（分）',
                                  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                                  PRIMARY KEY (`id`),
                                  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_drug_name` (`drug_name`),
  KEY `idx_expiry_date` (`expiry_date`),
  CONSTRAINT `fk_drug_inventory_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='药品库存表';

CREATE TABLE `drug_transaction` (
                                    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                    `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                    `drug_id` bigint NOT NULL COMMENT '药品ID',
                                    `type` tinyint NOT NULL COMMENT '类型：1-入库，2-出库',
                                    `quantity` int NOT NULL COMMENT '数量',
                                    `transaction_time` datetime NOT NULL COMMENT '操作时间',
                                    `operator_id` bigint DEFAULT NULL COMMENT '操作人ID',
                                    `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                    PRIMARY KEY (`id`),
                                    KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_drug_id` (`drug_id`),
  KEY `idx_transaction_time` (`transaction_time`),
  KEY `fk_drug_transactions_operator` (`operator_id`),
  CONSTRAINT `fk_drug_transactions_drug` FOREIGN KEY (`drug_id`) REFERENCES `drug_inventory` (`id`),
  CONSTRAINT `fk_drug_transactions_operator` FOREIGN KEY (`operator_id`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_drug_transactions_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='药品出入库记录表';

CREATE TABLE `elder` (
                         `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                         `tenant_id` bigint NOT NULL COMMENT '租户ID',
                         `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '姓名',
                         `id_card` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '身份证号（加密存储）',
                         `gender` tinyint DEFAULT NULL COMMENT '性别：0-女，1-男',
                         `birth_date` date DEFAULT NULL COMMENT '出生日期',
                         `phone` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '本人电话',
                         `emergency_contact` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '紧急联系人',
                         `emergency_phone` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '紧急联系电话',
                         `avatar` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像',
                         `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-咨询，1-试住，2-在住，3-退住，4-离世',
                         `admission_date` date DEFAULT NULL COMMENT '入住日期',
                         `care_level_id` bigint DEFAULT NULL COMMENT '当前护理等级ID',
                         `bed_id` bigint DEFAULT NULL COMMENT '当前床位ID',
                         `dietary_type` tinyint DEFAULT NULL COMMENT '饮食类型：1-普食，2-流食，3-半流食，4-素食',
                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                         `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                         PRIMARY KEY (`id`),
                         KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_care_level_id` (`care_level_id`),
  KEY `idx_bed_id` (`bed_id`),
  KEY `idx_id_card` (`id_card`),
  CONSTRAINT `fk_elders_bed` FOREIGN KEY (`bed_id`) REFERENCES `bed` (`id`),
  CONSTRAINT `fk_elders_care_level` FOREIGN KEY (`care_level_id`) REFERENCES `care_level` (`id`),
  CONSTRAINT `fk_elders_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='老人档案表';

CREATE TABLE `evaluation` (
                              `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                              `tenant_id` bigint NOT NULL COMMENT '租户ID',
                              `elderly_id` bigint NOT NULL COMMENT '老人ID',
                              `target_type` tinyint NOT NULL COMMENT '评价对象类型：1-护理员，2-餐饮，3-环境',
                              `target_id` bigint NOT NULL COMMENT '具体目标ID',
                              `score` tinyint NOT NULL COMMENT '评分1-5',
                              `comment` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '评论',
                              `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              PRIMARY KEY (`id`),
                              KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  KEY `idx_target` (`target_type`,`target_id`),
  CONSTRAINT `fk_evaluations_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_evaluations_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='服务评价表';

CREATE TABLE `fee_item` (
                            `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                            `tenant_id` bigint NOT NULL COMMENT '租户ID',
                            `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '费用项名称',
                            `fee_type` tinyint NOT NULL COMMENT '类型：1-一次性，2-周期性，3-按次',
                            `calculation_type` tinyint NOT NULL COMMENT '计算方式：1-固定金额，2-按天，3-按用量',
                            `default_price` bigint NOT NULL DEFAULT '0' COMMENT '默认单价（分）',
                            `billing_cycle` tinyint DEFAULT NULL COMMENT '周期性费用计费周期：1-月，2-季，3-年',
                            `is_active` tinyint NOT NULL DEFAULT '1' COMMENT '是否启用：0-否，1-是',
                            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                            PRIMARY KEY (`id`),
                            KEY `idx_tenant_id` (`tenant_id`),
  CONSTRAINT `fk_fee_items_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='费用项配置表';

CREATE TABLE `handover` (
                            `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                            `tenant_id` bigint NOT NULL COMMENT '租户ID',
                            `from_staff_id` bigint NOT NULL COMMENT '交班员工ID',
                            `to_staff_id` bigint NOT NULL COMMENT '接班员工ID',
                            `shift_start` datetime NOT NULL COMMENT '班次开始',
                            `shift_end` datetime NOT NULL COMMENT '班次结束',
                            `content` text COLLATE utf8mb4_unicode_ci COMMENT '交接内容',
                            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            PRIMARY KEY (`id`),
                            KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_from_staff_id` (`from_staff_id`),
  KEY `idx_to_staff_id` (`to_staff_id`),
  CONSTRAINT `fk_handovers_from_staff` FOREIGN KEY (`from_staff_id`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_handovers_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`),
  CONSTRAINT `fk_handovers_to_staff` FOREIGN KEY (`to_staff_id`) REFERENCES `staff` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='交接班记录表';

CREATE TABLE `health_alert` (
                                `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                `elderly_id` bigint NOT NULL COMMENT '老人ID',
                                `alert_type` tinyint NOT NULL COMMENT '预警类型：1-体征异常，2-用药逾期，3-跌倒报警',
                                `content` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '预警内容',
                                `alert_time` datetime NOT NULL COMMENT '预警时间',
                                `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-未处理，1-已处理',
                                `handler_id` bigint DEFAULT NULL COMMENT '处理人ID',
                                `handled_at` datetime DEFAULT NULL COMMENT '处理时间',
                                `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                PRIMARY KEY (`id`),
                                KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  KEY `idx_status` (`status`),
  KEY `fk_health_alerts_handler` (`handler_id`),
  CONSTRAINT `fk_health_alerts_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_health_alerts_handler` FOREIGN KEY (`handler_id`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_health_alerts_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='健康预警表';

CREATE TABLE `health_record` (
                                 `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                 `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                 `elderly_id` bigint NOT NULL COMMENT '老人ID',
                                 `past_medical_history` text COLLATE utf8mb4_unicode_ci COMMENT '既往病史',
                                 `allergy` text COLLATE utf8mb4_unicode_ci COMMENT '过敏史',
                                 `family_history` text COLLATE utf8mb4_unicode_ci COMMENT '家族史',
                                 `blood_type` varchar(5) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '血型',
                                 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                 PRIMARY KEY (`id`),
                                 KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  CONSTRAINT `fk_health_records_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_health_records_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='健康档案表';

CREATE TABLE `institutions` (
                                `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '院区名称',
                                `address` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '地址',
                                `phone` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '联系电话',
                                `principal` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '负责人',
                                `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                                PRIMARY KEY (`id`),
                                KEY `idx_tenant_id` (`tenant_id`),
  CONSTRAINT `fk_institutions_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='机构信息表';

CREATE TABLE `iot_data` (
                            `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                            `tenant_id` bigint NOT NULL COMMENT '租户ID',
                            `device_id` bigint NOT NULL COMMENT '设备ID',
                            `data_type` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '数据类型',
                            `data_value` json DEFAULT NULL COMMENT '数据值',
                            `event_time` datetime NOT NULL COMMENT '事件时间',
                            `is_alert` tinyint NOT NULL DEFAULT '0' COMMENT '是否触发告警：0-否，1-是',
                            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            PRIMARY KEY (`id`),
                            KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_device_id` (`device_id`),
  KEY `idx_event_time` (`event_time`),
  CONSTRAINT `fk_iot_data_device` FOREIGN KEY (`device_id`) REFERENCES `iot_device` (`id`),
  CONSTRAINT `fk_iot_data_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='设备数据记录表';

CREATE TABLE `iot_device` (
                              `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                              `tenant_id` bigint NOT NULL COMMENT '租户ID',
                              `device_id` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '设备唯一标识',
                              `device_type` tinyint NOT NULL COMMENT '设备类型：1-智能床垫，2-手环，3-跌倒雷达，4-门禁',
                              `elderly_id` bigint DEFAULT NULL COMMENT '绑定老人ID',
                              `location` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '安装位置',
                              `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-离线，1-在线',
                              `last_online` datetime DEFAULT NULL COMMENT '最后在线时间',
                              `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                              `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `uk_device_id` (`device_id`),
                              KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  CONSTRAINT `fk_iot_devices_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_iot_devices_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='设备表';

CREATE TABLE `material` (
                            `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                            `tenant_id` bigint NOT NULL COMMENT '租户ID',
                            `category_id` bigint DEFAULT NULL COMMENT '类别ID',
                            `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '物资名称',
                            `unit` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '单位',
                            `quantity` int NOT NULL DEFAULT '0' COMMENT '当前库存',
                            `price` bigint DEFAULT '0' COMMENT '单价（分）',
                            `location` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '存放位置',
                            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                            PRIMARY KEY (`id`),
                            KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_category_id` (`category_id`),
  CONSTRAINT `fk_materials_category` FOREIGN KEY (`category_id`) REFERENCES `material_categorie` (`id`),
  CONSTRAINT `fk_materials_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='物资库存表';

CREATE TABLE `material_categorie` (
                                      `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                      `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                      `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '类别名称',
                                      `parent_id` bigint DEFAULT NULL COMMENT '上级类别ID',
                                      `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                      `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                      `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                                      PRIMARY KEY (`id`),
                                      KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_parent_id` (`parent_id`),
  CONSTRAINT `fk_material_categories_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='物资类别表';

CREATE TABLE `material_usage` (
                                  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                  `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                  `material_id` bigint NOT NULL COMMENT '物资ID',
                                  `elderly_id` bigint DEFAULT NULL COMMENT '领用老人ID',
                                  `quantity` int NOT NULL COMMENT '数量',
                                  `usage_date` datetime NOT NULL COMMENT '领用时间',
                                  `recipient` bigint DEFAULT NULL COMMENT '领用人（员工）ID',
                                  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  PRIMARY KEY (`id`),
                                  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_material_id` (`material_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  KEY `idx_recipient` (`recipient`),
  CONSTRAINT `fk_material_usages_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_material_usages_material` FOREIGN KEY (`material_id`) REFERENCES `material` (`id`),
  CONSTRAINT `fk_material_usages_staff` FOREIGN KEY (`recipient`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_material_usages_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='物资领用表';

CREATE TABLE `meal_deliverie` (
                                  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                  `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                  `meal_order_id` bigint NOT NULL COMMENT '关联订餐记录ID',
                                  `delivery_time` datetime DEFAULT NULL COMMENT '送达时间',
                                  `deliverer_id` bigint DEFAULT NULL COMMENT '送餐人ID',
                                  `photo_url` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '送达照片',
                                  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  PRIMARY KEY (`id`),
                                  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_meal_order_id` (`meal_order_id`),
  KEY `idx_deliverer_id` (`deliverer_id`),
  CONSTRAINT `fk_meal_deliveries_order` FOREIGN KEY (`meal_order_id`) REFERENCES `meal_order` (`id`),
  CONSTRAINT `fk_meal_deliveries_staff` FOREIGN KEY (`deliverer_id`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_meal_deliveries_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='送餐记录表';

CREATE TABLE `meal_order` (
                              `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                              `tenant_id` bigint NOT NULL COMMENT '租户ID',
                              `elderly_id` bigint NOT NULL COMMENT '老人ID',
                              `order_date` date NOT NULL COMMENT '订餐日期',
                              `meal_type` tinyint NOT NULL COMMENT '餐次：1-早餐，2-午餐，3-晚餐，4-加餐',
                              `dietary_type` tinyint NOT NULL COMMENT '饮食类型',
                              `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-已预订，1-已配送，2-已取消',
                              `order_source` tinyint NOT NULL COMMENT '来源：1-家属端，2-护理员代订',
                              `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                              PRIMARY KEY (`id`),
                              KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  KEY `idx_order_date` (`order_date`),
  CONSTRAINT `fk_meal_orders_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_meal_orders_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订餐记录表';

CREATE TABLE `medication_record` (
                                     `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                     `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                     `elderly_id` bigint NOT NULL COMMENT '老人ID',
                                     `prescription_id` bigint DEFAULT NULL COMMENT '关联处方ID',
                                     `scheduled_time` datetime NOT NULL COMMENT '计划用药时间',
                                     `actual_time` datetime DEFAULT NULL COMMENT '实际用药时间',
                                     `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-未服，1-已服，2-拒服，3-漏服',
                                     `executor_id` bigint DEFAULT NULL COMMENT '执行护理员ID',
                                     `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                     `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                     `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                     PRIMARY KEY (`id`),
                                     KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  KEY `idx_prescription_id` (`prescription_id`),
  KEY `idx_scheduled_time` (`scheduled_time`),
  KEY `idx_executor_id` (`executor_id`),
  CONSTRAINT `fk_medication_records_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_medication_records_executor` FOREIGN KEY (`executor_id`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_medication_records_prescription` FOREIGN KEY (`prescription_id`) REFERENCES `prescription` (`id`),
  CONSTRAINT `fk_medication_records_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用药记录表';

CREATE TABLE `message` (
                           `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                           `tenant_id` bigint NOT NULL COMMENT '租户ID',
                           `elderly_id` bigint NOT NULL COMMENT '老人ID',
                           `sender_type` tinyint NOT NULL COMMENT '发送者类型：1-家属，2-护理员',
                           `sender_id` bigint NOT NULL COMMENT '发送人ID',
                           `content` text COLLATE utf8mb4_unicode_ci COMMENT '内容',
                           `photo_url` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '图片',
                           `parent_id` bigint DEFAULT NULL COMMENT '回复的留言ID',
                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           PRIMARY KEY (`id`),
                           KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  KEY `idx_parent_id` (`parent_id`),
  CONSTRAINT `fk_messages_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_messages_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='家属留言表';

CREATE TABLE `notification` (
                                `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                `user_id` bigint NOT NULL COMMENT '接收人ID',
                                `user_type` tinyint NOT NULL COMMENT '接收人类型：1-员工，2-家属',
                                `title` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标题',
                                `content` text COLLATE utf8mb4_unicode_ci COMMENT '内容',
                                `type` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '消息类型',
                                `is_read` tinyint NOT NULL DEFAULT '0' COMMENT '是否已读：0-未读，1-已读',
                                `sent_at` datetime NOT NULL COMMENT '发送时间',
                                `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                PRIMARY KEY (`id`),
                                KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_is_read` (`is_read`),
  CONSTRAINT `fk_notifications_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='消息通知表';

CREATE TABLE `outgoing_record` (
                                   `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                   `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                   `elderly_id` bigint NOT NULL COMMENT '老人ID',
                                   `type` tinyint NOT NULL COMMENT '类型：1-请假外出，2-临时外出',
                                   `start_time` datetime NOT NULL COMMENT '外出开始时间',
                                   `end_time` datetime DEFAULT NULL COMMENT '计划返回时间',
                                   `actual_return_time` datetime DEFAULT NULL COMMENT '实际返回时间',
                                   `reason` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '原因',
                                   `companion` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '陪同人',
                                   `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-外出中，1-已返回，2-超时未归',
                                   `approved_by` bigint DEFAULT NULL COMMENT '审批人ID',
                                   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                   PRIMARY KEY (`id`),
                                   KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  KEY `idx_status` (`status`),
  KEY `fk_outgoing_records_approver` (`approved_by`),
  CONSTRAINT `fk_outgoing_records_approver` FOREIGN KEY (`approved_by`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_outgoing_records_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_outgoing_records_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='出入登记表';

CREATE TABLE `package` (
                           `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                           `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '套餐名称',
                           `price` bigint NOT NULL DEFAULT '0' COMMENT '价格（分）',
                           `billing_cycle` tinyint NOT NULL DEFAULT '1' COMMENT '计费周期：1-月，2-年',
                           `max_staff` int NOT NULL DEFAULT '-1' COMMENT '最大员工数（-1不限）',
                           `max_beds` int NOT NULL DEFAULT '-1' COMMENT '最大床位数（-1不限）',
                           `features` json DEFAULT NULL COMMENT '功能权限配置',
                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='套餐表';

CREATE TABLE `patrol_point` (
                                `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '巡更点名称',
                                `location` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '位置描述',
                                `qr_code` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '二维码标识',
                                `sort_order` int DEFAULT '0' COMMENT '巡更顺序',
                                `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                                PRIMARY KEY (`id`),
                                KEY `idx_tenant_id` (`tenant_id`),
  CONSTRAINT `fk_patrol_points_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='巡更点表';

CREATE TABLE `patrol_record` (
                                 `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                 `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                 `patrol_point_id` bigint NOT NULL COMMENT '巡更点ID',
                                 `staff_id` bigint NOT NULL COMMENT '巡更员工ID',
                                 `patrol_time` datetime NOT NULL COMMENT '巡更时间',
                                 `photo_url` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '现场照片',
                                 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 PRIMARY KEY (`id`),
                                 KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_patrol_point_id` (`patrol_point_id`),
  KEY `idx_staff_id` (`staff_id`),
  CONSTRAINT `fk_patrol_records_point` FOREIGN KEY (`patrol_point_id`) REFERENCES `patrol_point` (`id`),
  CONSTRAINT `fk_patrol_records_staff` FOREIGN KEY (`staff_id`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_patrol_records_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='巡更记录表';

CREATE TABLE `payment` (
                           `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                           `tenant_id` bigint NOT NULL COMMENT '租户ID',
                           `elderly_id` bigint NOT NULL COMMENT '老人ID',
                           `bill_id` bigint DEFAULT NULL COMMENT '关联账单ID',
                           `amount` bigint NOT NULL DEFAULT '0' COMMENT '缴费金额（分）',
                           `payment_method` tinyint NOT NULL COMMENT '支付方式：1-现金，2-微信，3-支付宝，4-银行转账',
                           `transaction_id` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '第三方交易号',
                           `paid_at` datetime NOT NULL COMMENT '支付时间',
                           `operator_id` bigint DEFAULT NULL COMMENT '收款人ID',
                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           PRIMARY KEY (`id`),
                           KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  KEY `idx_bill_id` (`bill_id`),
  KEY `idx_transaction_id` (`transaction_id`),
  KEY `fk_payments_operator` (`operator_id`),
  CONSTRAINT `fk_payments_bill` FOREIGN KEY (`bill_id`) REFERENCES `bill` (`id`),
  CONSTRAINT `fk_payments_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_payments_operator` FOREIGN KEY (`operator_id`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_payments_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='缴费记录表';

CREATE TABLE `prescription` (
                                `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                `elderly_id` bigint NOT NULL COMMENT '老人ID',
                                `drug_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '药品名称',
                                `dosage` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '每次剂量',
                                `frequency` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用法频率',
                                `start_date` date NOT NULL COMMENT '开始日期',
                                `end_date` date DEFAULT NULL COMMENT '结束日期',
                                `doctor` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '开方医生',
                                `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：1-进行中，0-已停止',
                                `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                                PRIMARY KEY (`id`),
                                KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  CONSTRAINT `fk_prescriptions_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_prescriptions_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='处方表';

CREATE TABLE `recipes` (
                           `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                           `tenant_id` bigint NOT NULL COMMENT '租户ID',
                           `date` date NOT NULL COMMENT '日期',
                           `meal_type` tinyint NOT NULL COMMENT '餐次：1-早餐，2-午餐，3-晚餐，4-加餐',
                           `dietary_type` tinyint NOT NULL COMMENT '饮食类型：1-普食，2-流食，3-半流食，4-素食',
                           `content` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '食谱内容',
                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           PRIMARY KEY (`id`),
                           KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_date` (`date`),
  CONSTRAINT `fk_recipes_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='食谱表';

CREATE TABLE `repair` (
                          `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                          `tenant_id` bigint NOT NULL COMMENT '租户ID',
                          `location` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '报修位置',
                          `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '故障描述',
                          `reporter_id` bigint DEFAULT NULL COMMENT '报修人ID',
                          `report_time` datetime NOT NULL COMMENT '报修时间',
                          `assigned_to` bigint DEFAULT NULL COMMENT '维修人ID',
                          `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-待派单，1-维修中，2-已完成，3-已关闭',
                          `completion_time` datetime DEFAULT NULL COMMENT '完成时间',
                          `evaluation` tinyint DEFAULT NULL COMMENT '满意度评分1-5',
                          `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                          PRIMARY KEY (`id`),
                          KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_reporter_id` (`reporter_id`),
  KEY `idx_assigned_to` (`assigned_to`),
  KEY `idx_status` (`status`),
  CONSTRAINT `fk_repairs_assignee` FOREIGN KEY (`assigned_to`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_repairs_reporter` FOREIGN KEY (`reporter_id`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_repairs_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='报修记录表';

CREATE TABLE `role` (
                        `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                        `tenant_id` bigint NOT NULL COMMENT '租户ID（平台角色tenant_id=0）',
                        `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名称',
                        `permissions` json DEFAULT NULL COMMENT '权限列表',
                        `is_system` tinyint NOT NULL DEFAULT '0' COMMENT '是否系统预置：0-否，1-是',
                        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                        PRIMARY KEY (`id`),
                        KEY `idx_tenant_id` (`tenant_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

CREATE TABLE `room` (
                        `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                        `tenant_id` bigint NOT NULL COMMENT '租户ID',
                        `institution_id` bigint NOT NULL COMMENT '所属院区ID',
                        `building` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '楼栋',
                        `floor` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '楼层',
                        `room_no` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '房间号',
                        `type` tinyint DEFAULT '1' COMMENT '房间类型：1-单人间，2-双人间，3-多人间',
                        `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-空闲，1-部分占用，2-全满，3-维修',
                        `price` bigint DEFAULT '0' COMMENT '房间基准价（分）',
                        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                        PRIMARY KEY (`id`),
                        KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_institution_id` (`institution_id`),
  CONSTRAINT `fk_rooms_institution` FOREIGN KEY (`institution_id`) REFERENCES `institutions` (`id`),
  CONSTRAINT `fk_rooms_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='房间表';

CREATE TABLE `room_transfer` (
                                 `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                 `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                 `elderly_id` bigint NOT NULL COMMENT '老人ID',
                                 `from_bed_id` bigint NOT NULL COMMENT '原床位ID',
                                 `to_bed_id` bigint NOT NULL COMMENT '新床位ID',
                                 `transfer_date` date NOT NULL COMMENT '转房日期',
                                 `reason` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '原因',
                                 `operator_id` bigint DEFAULT NULL COMMENT '操作人ID',
                                 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                 PRIMARY KEY (`id`),
                                 KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  KEY `idx_from_bed_id` (`from_bed_id`),
  KEY `idx_to_bed_id` (`to_bed_id`),
  CONSTRAINT `fk_room_transfers_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_room_transfers_from_bed` FOREIGN KEY (`from_bed_id`) REFERENCES `bed` (`id`),
  CONSTRAINT `fk_room_transfers_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`),
  CONSTRAINT `fk_room_transfers_to_bed` FOREIGN KEY (`to_bed_id`) REFERENCES `bed` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='转房记录表';

CREATE TABLE `staff` (
                         `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                         `tenant_id` bigint NOT NULL COMMENT '租户ID',
                         `institution_id` bigint DEFAULT NULL COMMENT '所属院区ID',
                         `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '姓名',
                         `phone` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '手机号（登录账号）',
                         `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '加密密码',
                         `role_id` bigint DEFAULT NULL COMMENT '角色ID',
                         `position` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '岗位',
                         `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：0-离职，1-在职',
                         `avatar` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像',
                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                         `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                         PRIMARY KEY (`id`),
                         KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_institution_id` (`institution_id`),
  KEY `idx_role_id` (`role_id`),
  KEY `idx_phone` (`phone`),
  CONSTRAINT `fk_staff_institution` FOREIGN KEY (`institution_id`) REFERENCES `institutions` (`id`),
  CONSTRAINT `fk_staff_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`),
  CONSTRAINT `fk_staff_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='员工表';

CREATE TABLE `staff_performance` (
                                     `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                     `tenant_id` bigint NOT NULL COMMENT '租户ID',
                                     `staff_id` bigint NOT NULL COMMENT '员工ID',
                                     `period` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '绩效周期（年月）',
                                     `task_count` int NOT NULL DEFAULT '0' COMMENT '完成任务数',
                                     `evaluation_score` decimal(3,2) DEFAULT NULL COMMENT '服务评价平均分',
                                     `bonus` bigint NOT NULL DEFAULT '0' COMMENT '绩效奖金（分）',
                                     `calculated_at` datetime DEFAULT NULL COMMENT '计算时间',
                                     `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                     `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                     PRIMARY KEY (`id`),
                                     KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_staff_id` (`staff_id`),
  KEY `idx_period` (`period`),
  CONSTRAINT `fk_staff_performances_staff` FOREIGN KEY (`staff_id`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_staff_performances_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='员工绩效表';

CREATE TABLE `tenant` (
                          `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                          `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '租户名称',
                          `code` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '租户编码',
                          `package_id` bigint DEFAULT NULL COMMENT '套餐ID',
                          `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：0-停用，1-启用',
                          `expire_at` datetime DEFAULT NULL COMMENT '套餐到期时间',
                          `contact_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '联系人',
                          `contact_phone` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '联系电话',
                          `logo_url` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Logo地址',
                          `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                          `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `uk_code` (`code`),
                          KEY `idx_package_id` (`package_id`),
  CONSTRAINT `fk_tenants_package` FOREIGN KEY (`package_id`) REFERENCES `package` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='租户表';

CREATE TABLE `visitor` (
                           `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                           `tenant_id` bigint NOT NULL COMMENT '租户ID',
                           `elderly_id` bigint NOT NULL COMMENT '被访老人ID',
                           `visitor_name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '访客姓名',
                           `visitor_id_card` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '身份证号（加密）',
                           `visit_time` datetime DEFAULT NULL COMMENT '进入时间',
                           `leave_time` datetime DEFAULT NULL COMMENT '离开时间',
                           `purpose` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '探访目的',
                           `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-预约，1-已进入，2-已离开',
                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           PRIMARY KEY (`id`),
                           KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  CONSTRAINT `fk_visitors_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_visitors_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='访客记录表';

CREATE TABLE `vital_sign` (
                              `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                              `tenant_id` bigint NOT NULL COMMENT '租户ID',
                              `elderly_id` bigint NOT NULL COMMENT '老人ID',
                              `record_time` datetime NOT NULL COMMENT '测量时间',
                              `bp_systolic` int DEFAULT NULL COMMENT '收缩压',
                              `bp_diastolic` int DEFAULT NULL COMMENT '舒张压',
                              `heart_rate` int DEFAULT NULL COMMENT '心率',
                              `temperature` decimal(3,1) DEFAULT NULL COMMENT '体温',
                              `blood_sugar` decimal(4,1) DEFAULT NULL COMMENT '血糖',
                              `oxygen` int DEFAULT NULL COMMENT '血氧饱和度',
                              `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                              `recorder_id` bigint DEFAULT NULL COMMENT '记录人ID',
                              `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                              PRIMARY KEY (`id`),
                              KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elderly_id` (`elderly_id`),
  KEY `idx_record_time` (`record_time`),
  KEY `idx_recorder_id` (`recorder_id`),
  CONSTRAINT `fk_vital_signs_elderly` FOREIGN KEY (`elderly_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_vital_signs_recorder` FOREIGN KEY (`recorder_id`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_vital_signs_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='生命体征记录表';
