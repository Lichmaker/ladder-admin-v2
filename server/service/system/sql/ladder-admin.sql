
CREATE TABLE IF NOT EXISTS `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=2304 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  ladder_admin_v2_dev.casbin_rule 的数据：~317 rows (大约)
DELETE FROM `casbin_rule`;
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
	(2232, 'p', '887', '/announcement/createAnnouncement', 'POST', '', '', ''),
	(2233, 'p', '887', '/announcement/deleteAnnouncement', 'DELETE', '', '', ''),
	(2234, 'p', '887', '/announcement/deleteAnnouncementByIds', 'DELETE', '', '', ''),
	(2236, 'p', '887', '/announcement/findAnnouncement', 'GET', '', '', ''),
	(2237, 'p', '887', '/announcement/getAnnouncementList', 'GET', '', '', ''),
	(2235, 'p', '887', '/announcement/updateAnnouncement', 'PUT', '', '', ''),
	(2123, 'p', '887', '/api/createApi', 'POST', '', '', ''),
	(2124, 'p', '887', '/api/deleteApi', 'POST', '', '', ''),
	(2129, 'p', '887', '/api/deleteApisByIds', 'DELETE', '', '', ''),
	(2127, 'p', '887', '/api/getAllApis', 'POST', '', '', ''),
	(2128, 'p', '887', '/api/getApiById', 'POST', '', '', ''),
	(2126, 'p', '887', '/api/getApiList', 'POST', '', '', ''),
	(2125, 'p', '887', '/api/updateApi', 'POST', '', '', ''),
	(2130, 'p', '887', '/authority/copyAuthority', 'POST', '', '', ''),
	(2131, 'p', '887', '/authority/createAuthority', 'POST', '', '', ''),
	(2132, 'p', '887', '/authority/deleteAuthority', 'POST', '', '', ''),
	(2134, 'p', '887', '/authority/getAuthorityList', 'POST', '', '', ''),
	(2135, 'p', '887', '/authority/setDataAuthority', 'POST', '', '', ''),
	(2133, 'p', '887', '/authority/updateAuthority', 'PUT', '', '', ''),
	(2200, 'p', '887', '/authorityBtn/canRemoveAuthorityBtn', 'POST', '', '', ''),
	(2199, 'p', '887', '/authorityBtn/getAuthorityBtn', 'POST', '', '', ''),
	(2198, 'p', '887', '/authorityBtn/setAuthorityBtn', 'POST', '', '', ''),
	(2171, 'p', '887', '/autoCode/createPackage', 'POST', '', '', ''),
	(2168, 'p', '887', '/autoCode/createPlug', 'POST', '', '', ''),
	(2165, 'p', '887', '/autoCode/createTemp', 'POST', '', '', ''),
	(2173, 'p', '887', '/autoCode/delPackage', 'POST', '', '', ''),
	(2177, 'p', '887', '/autoCode/delSysHistory', 'POST', '', '', ''),
	(2167, 'p', '887', '/autoCode/getColumn', 'GET', '', '', ''),
	(2163, 'p', '887', '/autoCode/getDB', 'GET', '', '', ''),
	(2174, 'p', '887', '/autoCode/getMeta', 'POST', '', '', ''),
	(2172, 'p', '887', '/autoCode/getPackage', 'POST', '', '', ''),
	(2176, 'p', '887', '/autoCode/getSysHistory', 'POST', '', '', ''),
	(2164, 'p', '887', '/autoCode/getTables', 'GET', '', '', ''),
	(2169, 'p', '887', '/autoCode/installPlugin', 'POST', '', '', ''),
	(2166, 'p', '887', '/autoCode/preview', 'POST', '', '', ''),
	(2170, 'p', '887', '/autoCode/pubPlug', 'POST', '', '', ''),
	(2175, 'p', '887', '/autoCode/rollback', 'POST', '', '', ''),
	(2137, 'p', '887', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', ''),
	(2136, 'p', '887', '/casbin/updateCasbin', 'POST', '', '', ''),
	(2160, 'p', '887', '/customer/customer', 'DELETE', '', '', ''),
	(2161, 'p', '887', '/customer/customer', 'GET', '', '', ''),
	(2159, 'p', '887', '/customer/customer', 'POST', '', '', ''),
	(2158, 'p', '887', '/customer/customer', 'PUT', '', '', ''),
	(2162, 'p', '887', '/customer/customerList', 'GET', '', '', ''),
	(2210, 'p', '887', '/dataStatistics/createDataStatistics', 'POST', '', '', ''),
	(2211, 'p', '887', '/dataStatistics/deleteDataStatistics', 'DELETE', '', '', ''),
	(2212, 'p', '887', '/dataStatistics/deleteDataStatisticsByIds', 'DELETE', '', '', ''),
	(2214, 'p', '887', '/dataStatistics/findDataStatistics', 'GET', '', '', ''),
	(2215, 'p', '887', '/dataStatistics/getDataStatisticsList', 'GET', '', '', ''),
	(2213, 'p', '887', '/dataStatistics/updateDataStatistics', 'PUT', '', '', ''),
	(2197, 'p', '887', '/email/emailSend', 'POST', '', '', ''),
	(2196, 'p', '887', '/email/emailTest', 'POST', '', '', ''),
	(2148, 'p', '887', '/fileUploadAndDownload/breakpointContinue', 'POST', '', '', ''),
	(2149, 'p', '887', '/fileUploadAndDownload/breakpointContinueFinish', 'POST', '', '', ''),
	(2152, 'p', '887', '/fileUploadAndDownload/deleteFile', 'POST', '', '', ''),
	(2153, 'p', '887', '/fileUploadAndDownload/editFileName', 'POST', '', '', ''),
	(2147, 'p', '887', '/fileUploadAndDownload/findFile', 'GET', '', '', ''),
	(2154, 'p', '887', '/fileUploadAndDownload/getFileList', 'POST', '', '', ''),
	(2150, 'p', '887', '/fileUploadAndDownload/removeChunk', 'POST', '', '', ''),
	(2151, 'p', '887', '/fileUploadAndDownload/upload', 'POST', '', '', ''),
	(2245, 'p', '887', '/inviteCode/createInviteCode', 'POST', '', '', ''),
	(2246, 'p', '887', '/inviteCode/deleteInviteCode', 'DELETE', '', '', ''),
	(2247, 'p', '887', '/inviteCode/deleteInviteCodeByIds', 'DELETE', '', '', ''),
	(2252, 'p', '887', '/inviteCode/deleteMyInviteCodeByIds', 'DELETE', '', '', ''),
	(2249, 'p', '887', '/inviteCode/findInviteCode', 'GET', '', '', ''),
	(2250, 'p', '887', '/inviteCode/getInviteCodeList', 'GET', '', '', ''),
	(2251, 'p', '887', '/inviteCode/getMyInviteCodeList', 'GET', '', '', ''),
	(2248, 'p', '887', '/inviteCode/updateInviteCode', 'PUT', '', '', ''),
	(2111, 'p', '887', '/jwt/jsonInBlacklist', 'POST', '', '', ''),
	(2138, 'p', '887', '/menu/addBaseMenu', 'POST', '', '', ''),
	(2146, 'p', '887', '/menu/addMenuAuthority', 'POST', '', '', ''),
	(2140, 'p', '887', '/menu/deleteBaseMenu', 'POST', '', '', ''),
	(2142, 'p', '887', '/menu/getBaseMenuById', 'POST', '', '', ''),
	(2144, 'p', '887', '/menu/getBaseMenuTree', 'POST', '', '', ''),
	(2139, 'p', '887', '/menu/getMenu', 'POST', '', '', ''),
	(2145, 'p', '887', '/menu/getMenuAuthority', 'POST', '', '', ''),
	(2143, 'p', '887', '/menu/getMenuList', 'POST', '', '', ''),
	(2141, 'p', '887', '/menu/updateBaseMenu', 'POST', '', '', ''),
	(2194, 'p', '887', '/simpleUploader/checkFileMd5', 'GET', '', '', ''),
	(2195, 'p', '887', '/simpleUploader/mergeFileMd5', 'GET', '', '', ''),
	(2193, 'p', '887', '/simpleUploader/upload', 'POST', '', '', ''),
	(2183, 'p', '887', '/sysDictionary/createSysDictionary', 'POST', '', '', ''),
	(2184, 'p', '887', '/sysDictionary/deleteSysDictionary', 'DELETE', '', '', ''),
	(2186, 'p', '887', '/sysDictionary/findSysDictionary', 'GET', '', '', ''),
	(2187, 'p', '887', '/sysDictionary/getSysDictionaryList', 'GET', '', '', ''),
	(2185, 'p', '887', '/sysDictionary/updateSysDictionary', 'PUT', '', '', ''),
	(2179, 'p', '887', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', '', '', ''),
	(2180, 'p', '887', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', '', '', ''),
	(2181, 'p', '887', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', ''),
	(2182, 'p', '887', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', ''),
	(2178, 'p', '887', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', '', '', ''),
	(2201, 'p', '887', '/sysExportTemplate/createSysExportTemplate', 'POST', '', '', ''),
	(2202, 'p', '887', '/sysExportTemplate/deleteSysExportTemplate', 'DELETE', '', '', ''),
	(2203, 'p', '887', '/sysExportTemplate/deleteSysExportTemplateByIds', 'DELETE', '', '', ''),
	(2207, 'p', '887', '/sysExportTemplate/exportExcel', 'GET', '', '', ''),
	(2208, 'p', '887', '/sysExportTemplate/exportTemplate', 'GET', '', '', ''),
	(2205, 'p', '887', '/sysExportTemplate/findSysExportTemplate', 'GET', '', '', ''),
	(2206, 'p', '887', '/sysExportTemplate/getSysExportTemplateList', 'GET', '', '', ''),
	(2209, 'p', '887', '/sysExportTemplate/importExcel', 'POST', '', '', ''),
	(2204, 'p', '887', '/sysExportTemplate/updateSysExportTemplate', 'PUT', '', '', ''),
	(2188, 'p', '887', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', ''),
	(2191, 'p', '887', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', '', '', ''),
	(2192, 'p', '887', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', '', '', ''),
	(2189, 'p', '887', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', ''),
	(2190, 'p', '887', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', ''),
	(2155, 'p', '887', '/system/getServerInfo', 'POST', '', '', ''),
	(2156, 'p', '887', '/system/getSystemConfig', 'POST', '', '', ''),
	(2157, 'p', '887', '/system/setSystemConfig', 'POST', '', '', ''),
	(2113, 'p', '887', '/user/admin_register', 'POST', '', '', ''),
	(2122, 'p', '887', '/user/batchQuery', 'POST', '', '', ''),
	(2119, 'p', '887', '/user/changePassword', 'POST', '', '', ''),
	(2112, 'p', '887', '/user/deleteUser', 'DELETE', '', '', ''),
	(2117, 'p', '887', '/user/getUserInfo', 'GET', '', '', ''),
	(2114, 'p', '887', '/user/getUserList', 'POST', '', '', ''),
	(2121, 'p', '887', '/user/resetPassword', 'POST', '', '', ''),
	(2116, 'p', '887', '/user/setSelfInfo', 'PUT', '', '', ''),
	(2118, 'p', '887', '/user/setUserAuthorities', 'POST', '', '', ''),
	(2120, 'p', '887', '/user/setUserAuthority', 'POST', '', '', ''),
	(2115, 'p', '887', '/user/setUserInfo', 'PUT', '', '', ''),
	(2226, 'p', '887', '/userDataUsage/createUserDataUsage', 'POST', '', '', ''),
	(2227, 'p', '887', '/userDataUsage/deleteUserDataUsage', 'DELETE', '', '', ''),
	(2228, 'p', '887', '/userDataUsage/deleteUserDataUsageByIds', 'DELETE', '', '', ''),
	(2230, 'p', '887', '/userDataUsage/findUserDataUsage', 'GET', '', '', ''),
	(2231, 'p', '887', '/userDataUsage/getUserDataUsageList', 'GET', '', '', ''),
	(2229, 'p', '887', '/userDataUsage/updateUserDataUsage', 'PUT', '', '', ''),
	(2216, 'p', '887', '/userExt/createUserExt', 'POST', '', '', ''),
	(2217, 'p', '887', '/userExt/deleteUserExt', 'DELETE', '', '', ''),
	(2218, 'p', '887', '/userExt/deleteUserExtByIds', 'DELETE', '', '', ''),
	(2220, 'p', '887', '/userExt/findUserExt', 'GET', '', '', ''),
	(2221, 'p', '887', '/userExt/getUserExtList', 'GET', '', '', ''),
	(2219, 'p', '887', '/userExt/updateUserExt', 'PUT', '', '', ''),
	(2223, 'p', '887', '/userExt/v2/create', 'POST', '', '', ''),
	(2225, 'p', '887', '/userExt/v2/dashboard', 'GET', '', '', ''),
	(2222, 'p', '887', '/userExt/v2/getList', 'GET', '', '', ''),
	(2224, 'p', '887', '/userExt/v2/update', 'PUT', '', '', ''),
	(2238, 'p', '887', '/v2rayNode/createV2rayNode', 'POST', '', '', ''),
	(2239, 'p', '887', '/v2rayNode/deleteV2rayNode', 'DELETE', '', '', ''),
	(2240, 'p', '887', '/v2rayNode/deleteV2rayNodeByIds', 'DELETE', '', '', ''),
	(2242, 'p', '887', '/v2rayNode/findV2rayNode', 'GET', '', '', ''),
	(2243, 'p', '887', '/v2rayNode/getV2rayNodeList', 'GET', '', '', ''),
	(2244, 'p', '887', '/v2rayNode/pushData', 'POST', '', '', ''),
	(2241, 'p', '887', '/v2rayNode/updateV2rayNode', 'PUT', '', '', ''),
	(2090, 'p', '888', '/announcement/createAnnouncement', 'POST', '', '', ''),
	(2091, 'p', '888', '/announcement/deleteAnnouncement', 'DELETE', '', '', ''),
	(2092, 'p', '888', '/announcement/deleteAnnouncementByIds', 'DELETE', '', '', ''),
	(2094, 'p', '888', '/announcement/findAnnouncement', 'GET', '', '', ''),
	(2095, 'p', '888', '/announcement/getAnnouncementList', 'GET', '', '', ''),
	(2093, 'p', '888', '/announcement/updateAnnouncement', 'PUT', '', '', ''),
	(1982, 'p', '888', '/api/createApi', 'POST', '', '', ''),
	(1983, 'p', '888', '/api/deleteApi', 'POST', '', '', ''),
	(1988, 'p', '888', '/api/deleteApisByIds', 'DELETE', '', '', ''),
	(1986, 'p', '888', '/api/getAllApis', 'POST', '', '', ''),
	(1987, 'p', '888', '/api/getApiById', 'POST', '', '', ''),
	(1985, 'p', '888', '/api/getApiList', 'POST', '', '', ''),
	(1984, 'p', '888', '/api/updateApi', 'POST', '', '', ''),
	(1989, 'p', '888', '/authority/copyAuthority', 'POST', '', '', ''),
	(1990, 'p', '888', '/authority/createAuthority', 'POST', '', '', ''),
	(1991, 'p', '888', '/authority/deleteAuthority', 'POST', '', '', ''),
	(1993, 'p', '888', '/authority/getAuthorityList', 'POST', '', '', ''),
	(1994, 'p', '888', '/authority/setDataAuthority', 'POST', '', '', ''),
	(1992, 'p', '888', '/authority/updateAuthority', 'PUT', '', '', ''),
	(2058, 'p', '888', '/authorityBtn/canRemoveAuthorityBtn', 'POST', '', '', ''),
	(2057, 'p', '888', '/authorityBtn/getAuthorityBtn', 'POST', '', '', ''),
	(2056, 'p', '888', '/authorityBtn/setAuthorityBtn', 'POST', '', '', ''),
	(2030, 'p', '888', '/autoCode/createPackage', 'POST', '', '', ''),
	(2027, 'p', '888', '/autoCode/createPlug', 'POST', '', '', ''),
	(2024, 'p', '888', '/autoCode/createTemp', 'POST', '', '', ''),
	(2032, 'p', '888', '/autoCode/delPackage', 'POST', '', '', ''),
	(2036, 'p', '888', '/autoCode/delSysHistory', 'POST', '', '', ''),
	(2026, 'p', '888', '/autoCode/getColumn', 'GET', '', '', ''),
	(2022, 'p', '888', '/autoCode/getDB', 'GET', '', '', ''),
	(2033, 'p', '888', '/autoCode/getMeta', 'POST', '', '', ''),
	(2031, 'p', '888', '/autoCode/getPackage', 'POST', '', '', ''),
	(2035, 'p', '888', '/autoCode/getSysHistory', 'POST', '', '', ''),
	(2023, 'p', '888', '/autoCode/getTables', 'GET', '', '', ''),
	(2028, 'p', '888', '/autoCode/installPlugin', 'POST', '', '', ''),
	(2025, 'p', '888', '/autoCode/preview', 'POST', '', '', ''),
	(2029, 'p', '888', '/autoCode/pubPlug', 'POST', '', '', ''),
	(2034, 'p', '888', '/autoCode/rollback', 'POST', '', '', ''),
	(1996, 'p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', ''),
	(1995, 'p', '888', '/casbin/updateCasbin', 'POST', '', '', ''),
	(2019, 'p', '888', '/customer/customer', 'DELETE', '', '', ''),
	(2020, 'p', '888', '/customer/customer', 'GET', '', '', ''),
	(2018, 'p', '888', '/customer/customer', 'POST', '', '', ''),
	(2017, 'p', '888', '/customer/customer', 'PUT', '', '', ''),
	(2021, 'p', '888', '/customer/customerList', 'GET', '', '', ''),
	(2068, 'p', '888', '/dataStatistics/createDataStatistics', 'POST', '', '', ''),
	(2069, 'p', '888', '/dataStatistics/deleteDataStatistics', 'DELETE', '', '', ''),
	(2070, 'p', '888', '/dataStatistics/deleteDataStatisticsByIds', 'DELETE', '', '', ''),
	(2072, 'p', '888', '/dataStatistics/findDataStatistics', 'GET', '', '', ''),
	(2073, 'p', '888', '/dataStatistics/getDataStatisticsList', 'GET', '', '', ''),
	(2071, 'p', '888', '/dataStatistics/updateDataStatistics', 'PUT', '', '', ''),
	(2055, 'p', '888', '/email/emailTest', 'POST', '', '', ''),
	(2007, 'p', '888', '/fileUploadAndDownload/breakpointContinue', 'POST', '', '', ''),
	(2008, 'p', '888', '/fileUploadAndDownload/breakpointContinueFinish', 'POST', '', '', ''),
	(2011, 'p', '888', '/fileUploadAndDownload/deleteFile', 'POST', '', '', ''),
	(2012, 'p', '888', '/fileUploadAndDownload/editFileName', 'POST', '', '', ''),
	(2006, 'p', '888', '/fileUploadAndDownload/findFile', 'GET', '', '', ''),
	(2013, 'p', '888', '/fileUploadAndDownload/getFileList', 'POST', '', '', ''),
	(2009, 'p', '888', '/fileUploadAndDownload/removeChunk', 'POST', '', '', ''),
	(2010, 'p', '888', '/fileUploadAndDownload/upload', 'POST', '', '', ''),
	(2103, 'p', '888', '/inviteCode/createInviteCode', 'POST', '', '', ''),
	(2104, 'p', '888', '/inviteCode/deleteInviteCode', 'DELETE', '', '', ''),
	(2105, 'p', '888', '/inviteCode/deleteInviteCodeByIds', 'DELETE', '', '', ''),
	(2110, 'p', '888', '/inviteCode/deleteMyInviteCodeByIds', 'DELETE', '', '', ''),
	(2107, 'p', '888', '/inviteCode/findInviteCode', 'GET', '', '', ''),
	(2108, 'p', '888', '/inviteCode/getInviteCodeList', 'GET', '', '', ''),
	(2109, 'p', '888', '/inviteCode/getMyInviteCodeList', 'GET', '', '', ''),
	(2106, 'p', '888', '/inviteCode/updateInviteCode', 'PUT', '', '', ''),
	(1970, 'p', '888', '/jwt/jsonInBlacklist', 'POST', '', '', ''),
	(1997, 'p', '888', '/menu/addBaseMenu', 'POST', '', '', ''),
	(2005, 'p', '888', '/menu/addMenuAuthority', 'POST', '', '', ''),
	(1999, 'p', '888', '/menu/deleteBaseMenu', 'POST', '', '', ''),
	(2001, 'p', '888', '/menu/getBaseMenuById', 'POST', '', '', ''),
	(2003, 'p', '888', '/menu/getBaseMenuTree', 'POST', '', '', ''),
	(1998, 'p', '888', '/menu/getMenu', 'POST', '', '', ''),
	(2004, 'p', '888', '/menu/getMenuAuthority', 'POST', '', '', ''),
	(2002, 'p', '888', '/menu/getMenuList', 'POST', '', '', ''),
	(2000, 'p', '888', '/menu/updateBaseMenu', 'POST', '', '', ''),
	(2053, 'p', '888', '/simpleUploader/checkFileMd5', 'GET', '', '', ''),
	(2054, 'p', '888', '/simpleUploader/mergeFileMd5', 'GET', '', '', ''),
	(2052, 'p', '888', '/simpleUploader/upload', 'POST', '', '', ''),
	(2042, 'p', '888', '/sysDictionary/createSysDictionary', 'POST', '', '', ''),
	(2043, 'p', '888', '/sysDictionary/deleteSysDictionary', 'DELETE', '', '', ''),
	(2045, 'p', '888', '/sysDictionary/findSysDictionary', 'GET', '', '', ''),
	(2046, 'p', '888', '/sysDictionary/getSysDictionaryList', 'GET', '', '', ''),
	(2044, 'p', '888', '/sysDictionary/updateSysDictionary', 'PUT', '', '', ''),
	(2038, 'p', '888', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', '', '', ''),
	(2039, 'p', '888', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', '', '', ''),
	(2040, 'p', '888', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', ''),
	(2041, 'p', '888', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', ''),
	(2037, 'p', '888', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', '', '', ''),
	(2059, 'p', '888', '/sysExportTemplate/createSysExportTemplate', 'POST', '', '', ''),
	(2060, 'p', '888', '/sysExportTemplate/deleteSysExportTemplate', 'DELETE', '', '', ''),
	(2061, 'p', '888', '/sysExportTemplate/deleteSysExportTemplateByIds', 'DELETE', '', '', ''),
	(2065, 'p', '888', '/sysExportTemplate/exportExcel', 'GET', '', '', ''),
	(2066, 'p', '888', '/sysExportTemplate/exportTemplate', 'GET', '', '', ''),
	(2063, 'p', '888', '/sysExportTemplate/findSysExportTemplate', 'GET', '', '', ''),
	(2064, 'p', '888', '/sysExportTemplate/getSysExportTemplateList', 'GET', '', '', ''),
	(2067, 'p', '888', '/sysExportTemplate/importExcel', 'POST', '', '', ''),
	(2062, 'p', '888', '/sysExportTemplate/updateSysExportTemplate', 'PUT', '', '', ''),
	(2047, 'p', '888', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', ''),
	(2050, 'p', '888', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', '', '', ''),
	(2051, 'p', '888', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', '', '', ''),
	(2048, 'p', '888', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', ''),
	(2049, 'p', '888', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', ''),
	(2014, 'p', '888', '/system/getServerInfo', 'POST', '', '', ''),
	(2015, 'p', '888', '/system/getSystemConfig', 'POST', '', '', ''),
	(2016, 'p', '888', '/system/setSystemConfig', 'POST', '', '', ''),
	(1972, 'p', '888', '/user/admin_register', 'POST', '', '', ''),
	(1981, 'p', '888', '/user/batchQuery', 'POST', '', '', ''),
	(1978, 'p', '888', '/user/changePassword', 'POST', '', '', ''),
	(1971, 'p', '888', '/user/deleteUser', 'DELETE', '', '', ''),
	(1976, 'p', '888', '/user/getUserInfo', 'GET', '', '', ''),
	(1973, 'p', '888', '/user/getUserList', 'POST', '', '', ''),
	(1980, 'p', '888', '/user/resetPassword', 'POST', '', '', ''),
	(1975, 'p', '888', '/user/setSelfInfo', 'PUT', '', '', ''),
	(1977, 'p', '888', '/user/setUserAuthorities', 'POST', '', '', ''),
	(1979, 'p', '888', '/user/setUserAuthority', 'POST', '', '', ''),
	(1974, 'p', '888', '/user/setUserInfo', 'PUT', '', '', ''),
	(2084, 'p', '888', '/userDataUsage/createUserDataUsage', 'POST', '', '', ''),
	(2085, 'p', '888', '/userDataUsage/deleteUserDataUsage', 'DELETE', '', '', ''),
	(2086, 'p', '888', '/userDataUsage/deleteUserDataUsageByIds', 'DELETE', '', '', ''),
	(2088, 'p', '888', '/userDataUsage/findUserDataUsage', 'GET', '', '', ''),
	(2089, 'p', '888', '/userDataUsage/getUserDataUsageList', 'GET', '', '', ''),
	(2087, 'p', '888', '/userDataUsage/updateUserDataUsage', 'PUT', '', '', ''),
	(2074, 'p', '888', '/userExt/createUserExt', 'POST', '', '', ''),
	(2075, 'p', '888', '/userExt/deleteUserExt', 'DELETE', '', '', ''),
	(2076, 'p', '888', '/userExt/deleteUserExtByIds', 'DELETE', '', '', ''),
	(2078, 'p', '888', '/userExt/findUserExt', 'GET', '', '', ''),
	(2079, 'p', '888', '/userExt/getUserExtList', 'GET', '', '', ''),
	(2077, 'p', '888', '/userExt/updateUserExt', 'PUT', '', '', ''),
	(2081, 'p', '888', '/userExt/v2/create', 'POST', '', '', ''),
	(2083, 'p', '888', '/userExt/v2/dashboard', 'GET', '', '', ''),
	(2080, 'p', '888', '/userExt/v2/getList', 'GET', '', '', ''),
	(2082, 'p', '888', '/userExt/v2/update', 'PUT', '', '', ''),
	(2096, 'p', '888', '/v2rayNode/createV2rayNode', 'POST', '', '', ''),
	(2097, 'p', '888', '/v2rayNode/deleteV2rayNode', 'DELETE', '', '', ''),
	(2098, 'p', '888', '/v2rayNode/deleteV2rayNodeByIds', 'DELETE', '', '', ''),
	(2100, 'p', '888', '/v2rayNode/findV2rayNode', 'GET', '', '', ''),
	(2101, 'p', '888', '/v2rayNode/getV2rayNodeList', 'GET', '', '', ''),
	(2102, 'p', '888', '/v2rayNode/pushData', 'POST', '', '', ''),
	(2099, 'p', '888', '/v2rayNode/updateV2rayNode', 'PUT', '', '', ''),
	(2299, 'p', '9528', '/announcement/findAnnouncement', 'GET', '', '', ''),
	(2300, 'p', '9528', '/announcement/getAnnouncementList', 'GET', '', '', ''),
	(2261, 'p', '9528', '/api/createApi', 'POST', '', '', ''),
	(2262, 'p', '9528', '/api/deleteApi', 'POST', '', '', ''),
	(2265, 'p', '9528', '/api/getAllApis', 'POST', '', '', ''),
	(2266, 'p', '9528', '/api/getApiById', 'POST', '', '', ''),
	(2264, 'p', '9528', '/api/getApiList', 'POST', '', '', ''),
	(2263, 'p', '9528', '/api/updateApi', 'POST', '', '', ''),
	(2267, 'p', '9528', '/authority/createAuthority', 'POST', '', '', ''),
	(2268, 'p', '9528', '/authority/deleteAuthority', 'POST', '', '', ''),
	(2269, 'p', '9528', '/authority/getAuthorityList', 'POST', '', '', ''),
	(2270, 'p', '9528', '/authority/setDataAuthority', 'POST', '', '', ''),
	(2293, 'p', '9528', '/autoCode/createTemp', 'POST', '', '', ''),
	(2272, 'p', '9528', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', ''),
	(2271, 'p', '9528', '/casbin/updateCasbin', 'POST', '', '', ''),
	(2290, 'p', '9528', '/customer/customer', 'DELETE', '', '', ''),
	(2291, 'p', '9528', '/customer/customer', 'GET', '', '', ''),
	(2289, 'p', '9528', '/customer/customer', 'POST', '', '', ''),
	(2288, 'p', '9528', '/customer/customer', 'PUT', '', '', ''),
	(2292, 'p', '9528', '/customer/customerList', 'GET', '', '', ''),
	(2294, 'p', '9528', '/dataStatistics/findDataStatistics', 'GET', '', '', ''),
	(2295, 'p', '9528', '/dataStatistics/getDataStatisticsList', 'GET', '', '', ''),
	(2283, 'p', '9528', '/fileUploadAndDownload/deleteFile', 'POST', '', '', ''),
	(2284, 'p', '9528', '/fileUploadAndDownload/editFileName', 'POST', '', '', ''),
	(2285, 'p', '9528', '/fileUploadAndDownload/getFileList', 'POST', '', '', ''),
	(2282, 'p', '9528', '/fileUploadAndDownload/upload', 'POST', '', '', ''),
	(2303, 'p', '9528', '/inviteCode/deleteMyInviteCodeByIds', 'DELETE', '', '', ''),
	(2301, 'p', '9528', '/inviteCode/findInviteCode', 'GET', '', '', ''),
	(2302, 'p', '9528', '/inviteCode/getMyInviteCodeList', 'GET', '', '', ''),
	(2253, 'p', '9528', '/jwt/jsonInBlacklist', 'POST', '', '', ''),
	(2273, 'p', '9528', '/menu/addBaseMenu', 'POST', '', '', ''),
	(2281, 'p', '9528', '/menu/addMenuAuthority', 'POST', '', '', ''),
	(2275, 'p', '9528', '/menu/deleteBaseMenu', 'POST', '', '', ''),
	(2277, 'p', '9528', '/menu/getBaseMenuById', 'POST', '', '', ''),
	(2279, 'p', '9528', '/menu/getBaseMenuTree', 'POST', '', '', ''),
	(2274, 'p', '9528', '/menu/getMenu', 'POST', '', '', ''),
	(2280, 'p', '9528', '/menu/getMenuAuthority', 'POST', '', '', ''),
	(2278, 'p', '9528', '/menu/getMenuList', 'POST', '', '', ''),
	(2276, 'p', '9528', '/menu/updateBaseMenu', 'POST', '', '', ''),
	(2286, 'p', '9528', '/system/getSystemConfig', 'POST', '', '', ''),
	(2287, 'p', '9528', '/system/setSystemConfig', 'POST', '', '', ''),
	(2254, 'p', '9528', '/user/admin_register', 'POST', '', '', ''),
	(2260, 'p', '9528', '/user/batchQuery', 'POST', '', '', ''),
	(2258, 'p', '9528', '/user/changePassword', 'POST', '', '', ''),
	(2257, 'p', '9528', '/user/getUserInfo', 'GET', '', '', ''),
	(2255, 'p', '9528', '/user/getUserList', 'POST', '', '', ''),
	(2256, 'p', '9528', '/user/setSelfInfo', 'PUT', '', '', ''),
	(2259, 'p', '9528', '/user/setUserAuthority', 'POST', '', '', ''),
	(2297, 'p', '9528', '/userDataUsage/findUserDataUsage', 'GET', '', '', ''),
	(2298, 'p', '9528', '/userDataUsage/getUserDataUsageList', 'GET', '', '', ''),
	(2296, 'p', '9528', '/userExt/v2/dashboard', 'GET', '', '', '');


CREATE TABLE IF NOT EXISTS `exa_customers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint unsigned DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` bigint unsigned DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_exa_customers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `exa_customers`;


CREATE TABLE IF NOT EXISTS `exa_files` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `file_md5` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `file_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `chunk_total` bigint DEFAULT NULL,
  `is_finish` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_files_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `exa_files`;


CREATE TABLE IF NOT EXISTS `exa_file_chunks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `exa_file_id` bigint unsigned DEFAULT NULL,
  `file_chunk_number` bigint DEFAULT NULL,
  `file_chunk_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_chunks_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `exa_file_chunks`;


CREATE TABLE IF NOT EXISTS `exa_file_upload_and_downloads` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `exa_file_upload_and_downloads`;
INSERT INTO `exa_file_upload_and_downloads` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `url`, `tag`, `key`) VALUES
	(1, '2024-01-18 16:24:31.582', '2024-01-18 16:24:31.582', NULL, '10.png', 'https://qmplusimg.henrongyi.top/gvalogo.png', 'png', '158787308910.png'),
	(2, '2024-01-18 16:24:31.582', '2024-01-18 16:24:31.582', NULL, 'logo.png', 'https://qmplusimg.henrongyi.top/1576554439myAvatar.png', 'png', '1587973709logo.png');


CREATE TABLE IF NOT EXISTS `jwt_blacklists` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `jwt_blacklists`;
INSERT INTO `jwt_blacklists` (`id`, `created_at`, `updated_at`, `deleted_at`, `jwt`) VALUES
	(16, '2024-03-10 17:13:06.148', '2024-03-10 17:13:06.148', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNDFiOGY2MzQtNWVmYy1jZTA1LWVkNzAtMjAzYmQyY2MwMDVjIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6ImFkbWlu55So5oi3IiwiQXV0aG9yaXR5SWQiOjg4OCwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InFtUGx1cyIsImF1ZCI6WyJHVkEiXSwiZXhwIjoxNzEwNjY2Njk5LCJuYmYiOjE3MTAwNjE4OTl9.9-JQAJcsFjQJv_mFyKjeYi3pJZ-VzVUl4-qtPgqq6KI');


CREATE TABLE IF NOT EXISTS `la_announcement` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_la_announcement_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `la_announcement`;
INSERT INTO `la_announcement` (`id`, `title`, `content`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(2, '欢迎使用Ladder-Admin', '<p>欢迎使用Ladder-Admin</p>', '2024-02-04 13:53:18.807', '2024-03-08 16:06:27.804', NULL);


CREATE TABLE IF NOT EXISTS `la_data_statistics` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `node_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `date` date NOT NULL,
  `value` bigint unsigned NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email_node_code_date` (`email`,`node_id`,`date`) USING BTREE,
  KEY `email_date` (`email`,`date`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `la_data_statistics`;


CREATE TABLE IF NOT EXISTS `la_user_data_usage` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_uuid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `standard_data` int NOT NULL,
  `usage` bigint NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `start_date_end_date` (`start_date`,`end_date`),
  KEY `idx_la_user_data_usage_user_uuid` (`user_uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `la_user_data_usage`;


CREATE TABLE IF NOT EXISTS `la_user_ext` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_uuid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `valid_start` date DEFAULT NULL,
  `valid_end` date DEFAULT NULL,
  `standard_data` int NOT NULL COMMENT '基础用量(MB)',
  `enable` tinyint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`user_uuid`),
  KEY `idx_la_user_ext_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `la_user_ext`;
INSERT INTO `la_user_ext` (`created_at`, `updated_at`, `deleted_at`, `user_uuid`, `valid_start`, `valid_end`, `standard_data`, `enable`) VALUES
	('2024-03-06 12:09:04.539', '2024-03-07 11:00:04.723', NULL, '41b8f634-5efc-ce05-ed70-203bd2cc005c', '2024-03-06', '2030-04-05', 1024, 0);


CREATE TABLE IF NOT EXISTS `la_v2ray_node` (
  `id` int NOT NULL AUTO_INCREMENT,
  `unique_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `config_raw` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `secret` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `secret_iv` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `vmess_port` int DEFAULT NULL,
  `rpc_port` int DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `last_contact_at` datetime(3) DEFAULT NULL,
  `host` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `vpsproxy_port` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_id` (`unique_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `la_v2ray_node`;


CREATE TABLE IF NOT EXISTS `sys_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=144 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `sys_apis`;
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES
	(1, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单(退出，必选)', 'jwt', 'POST'),
	(2, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/user/deleteUser', '删除用户', '系统用户', 'DELETE'),
	(3, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/user/admin_register', '用户注册', '系统用户', 'POST'),
	(4, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/user/getUserList', '获取用户列表', '系统用户', 'POST'),
	(5, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/user/setUserInfo', '设置用户信息', '系统用户', 'PUT'),
	(6, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/user/setSelfInfo', '设置自身信息(必选)', '系统用户', 'PUT'),
	(7, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/user/getUserInfo', '获取自身信息(必选)', '系统用户', 'GET'),
	(8, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/user/setUserAuthorities', '设置权限组', '系统用户', 'POST'),
	(9, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/user/changePassword', '修改密码（建议选择)', '系统用户', 'POST'),
	(10, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/user/setUserAuthority', '修改用户角色(必选)', '系统用户', 'POST'),
	(11, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/user/resetPassword', '重置用户密码', '系统用户', 'POST'),
	(12, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/api/createApi', '创建api', 'api', 'POST'),
	(13, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/api/deleteApi', '删除Api', 'api', 'POST'),
	(14, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/api/updateApi', '更新Api', 'api', 'POST'),
	(15, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/api/getApiList', '获取api列表', 'api', 'POST'),
	(16, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/api/getAllApis', '获取所有api', 'api', 'POST'),
	(17, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/api/getApiById', '获取api详细信息', 'api', 'POST'),
	(18, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/api/deleteApisByIds', '批量删除api', 'api', 'DELETE'),
	(19, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/authority/copyAuthority', '拷贝角色', '角色', 'POST'),
	(20, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/authority/createAuthority', '创建角色', '角色', 'POST'),
	(21, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/authority/deleteAuthority', '删除角色', '角色', 'POST'),
	(22, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/authority/updateAuthority', '更新角色信息', '角色', 'PUT'),
	(23, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/authority/getAuthorityList', '获取角色列表', '角色', 'POST'),
	(24, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/authority/setDataAuthority', '设置角色资源权限', '角色', 'POST'),
	(25, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST'),
	(26, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST'),
	(27, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/menu/addBaseMenu', '新增菜单', '菜单', 'POST'),
	(28, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/menu/getMenu', '获取菜单树(必选)', '菜单', 'POST'),
	(29, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/menu/deleteBaseMenu', '删除菜单', '菜单', 'POST'),
	(30, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/menu/updateBaseMenu', '更新菜单', '菜单', 'POST'),
	(31, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/menu/getBaseMenuById', '根据id获取菜单', '菜单', 'POST'),
	(32, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/menu/getMenuList', '分页获取基础menu列表', '菜单', 'POST'),
	(33, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/menu/getBaseMenuTree', '获取用户动态路由', '菜单', 'POST'),
	(34, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/menu/getMenuAuthority', '获取指定角色menu', '菜单', 'POST'),
	(35, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', '菜单', 'POST'),
	(36, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/fileUploadAndDownload/findFile', '寻找目标文件（秒传）', '分片上传', 'GET'),
	(37, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/fileUploadAndDownload/breakpointContinue', '断点续传', '分片上传', 'POST'),
	(38, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/fileUploadAndDownload/breakpointContinueFinish', '断点续传完成', '分片上传', 'POST'),
	(39, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/fileUploadAndDownload/removeChunk', '上传完成移除文件', '分片上传', 'POST'),
	(40, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/fileUploadAndDownload/upload', '文件上传示例', '文件上传与下载', 'POST'),
	(41, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/fileUploadAndDownload/deleteFile', '删除文件', '文件上传与下载', 'POST'),
	(42, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/fileUploadAndDownload/editFileName', '文件名或者备注编辑', '文件上传与下载', 'POST'),
	(43, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/fileUploadAndDownload/getFileList', '获取上传文件列表', '文件上传与下载', 'POST'),
	(44, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/system/getServerInfo', '获取服务器信息', '系统服务', 'POST'),
	(45, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/system/getSystemConfig', '获取配置文件内容', '系统服务', 'POST'),
	(46, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/system/setSystemConfig', '设置配置文件内容', '系统服务', 'POST'),
	(47, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/customer/customer', '更新客户', '客户', 'PUT'),
	(48, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/customer/customer', '创建客户', '客户', 'POST'),
	(49, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/customer/customer', '删除客户', '客户', 'DELETE'),
	(50, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/customer/customer', '获取单一客户', '客户', 'GET'),
	(51, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/customer/customerList', '获取客户列表', '客户', 'GET'),
	(52, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/autoCode/getDB', '获取所有数据库', '代码生成器', 'GET'),
	(53, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/autoCode/getTables', '获取数据库表', '代码生成器', 'GET'),
	(54, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/autoCode/createTemp', '自动化代码', '代码生成器', 'POST'),
	(55, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/autoCode/preview', '预览自动化代码', '代码生成器', 'POST'),
	(56, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/autoCode/getColumn', '获取所选table的所有字段', '代码生成器', 'GET'),
	(57, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/autoCode/createPlug', '自动创建插件包', '代码生成器', 'POST'),
	(58, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/autoCode/installPlugin', '安装插件', '代码生成器', 'POST'),
	(59, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/autoCode/pubPlug', '打包插件', '代码生成器', 'POST'),
	(60, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/autoCode/createPackage', '生成包(package)', '包（pkg）生成器', 'POST'),
	(61, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/autoCode/getPackage', '获取所有包(package)', '包（pkg）生成器', 'POST'),
	(62, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/autoCode/delPackage', '删除包(package)', '包（pkg）生成器', 'POST'),
	(63, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/autoCode/getMeta', '获取meta信息', '代码生成器历史', 'POST'),
	(64, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/autoCode/rollback', '回滚自动生成代码', '代码生成器历史', 'POST'),
	(65, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/autoCode/getSysHistory', '查询回滚记录', '代码生成器历史', 'POST'),
	(66, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/autoCode/delSysHistory', '删除回滚记录', '代码生成器历史', 'POST'),
	(67, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysDictionaryDetail/updateSysDictionaryDetail', '更新字典内容', '系统字典详情', 'PUT'),
	(68, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysDictionaryDetail/createSysDictionaryDetail', '新增字典内容', '系统字典详情', 'POST'),
	(69, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysDictionaryDetail/deleteSysDictionaryDetail', '删除字典内容', '系统字典详情', 'DELETE'),
	(70, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysDictionaryDetail/findSysDictionaryDetail', '根据ID获取字典内容', '系统字典详情', 'GET'),
	(71, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysDictionaryDetail/getSysDictionaryDetailList', '获取字典内容列表', '系统字典详情', 'GET'),
	(72, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysDictionary/createSysDictionary', '新增字典', '系统字典', 'POST'),
	(73, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysDictionary/deleteSysDictionary', '删除字典', '系统字典', 'DELETE'),
	(74, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysDictionary/updateSysDictionary', '更新字典', '系统字典', 'PUT'),
	(75, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysDictionary/findSysDictionary', '根据ID获取字典', '系统字典', 'GET'),
	(76, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysDictionary/getSysDictionaryList', '获取字典列表', '系统字典', 'GET'),
	(77, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', '操作记录', 'POST'),
	(78, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', '操作记录', 'GET'),
	(79, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', '操作记录', 'GET'),
	(80, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', '操作记录', 'DELETE'),
	(81, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', '操作记录', 'DELETE'),
	(82, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/simpleUploader/upload', '插件版分片上传', '断点续传(插件版)', 'POST'),
	(83, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/simpleUploader/checkFileMd5', '文件完整度验证', '断点续传(插件版)', 'GET'),
	(84, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/simpleUploader/mergeFileMd5', '上传完成合并文件', '断点续传(插件版)', 'GET'),
	(85, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/email/emailTest', '发送测试邮件', 'email', 'POST'),
	(86, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/email/emailSend', '发送邮件示例', 'email', 'POST'),
	(87, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/authorityBtn/setAuthorityBtn', '设置按钮权限', '按钮权限', 'POST'),
	(88, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/authorityBtn/getAuthorityBtn', '获取已有按钮权限', '按钮权限', 'POST'),
	(89, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/authorityBtn/canRemoveAuthorityBtn', '删除按钮', '按钮权限', 'POST'),
	(90, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysExportTemplate/createSysExportTemplate', '新增导出模板', '表格模板', 'POST'),
	(91, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysExportTemplate/deleteSysExportTemplate', '删除导出模板', '表格模板', 'DELETE'),
	(92, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysExportTemplate/deleteSysExportTemplateByIds', '批量删除导出模板', '表格模板', 'DELETE'),
	(93, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysExportTemplate/updateSysExportTemplate', '更新导出模板', '表格模板', 'PUT'),
	(94, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysExportTemplate/findSysExportTemplate', '根据ID获取导出模板', '表格模板', 'GET'),
	(95, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysExportTemplate/getSysExportTemplateList', '获取导出模板列表', '表格模板', 'GET'),
	(96, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysExportTemplate/exportExcel', '导出Excel', '表格模板', 'GET'),
	(97, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysExportTemplate/exportTemplate', '下载模板', '表格模板', 'GET'),
	(98, '2024-01-18 16:24:20.471', '2024-01-18 16:24:20.471', NULL, '/sysExportTemplate/importExcel', '导入Excel', '表格模板', 'POST'),
	(99, '2024-01-22 15:44:57.525', '2024-01-22 15:44:57.525', NULL, '/dataStatistics/createDataStatistics', '新增dataStatistics表', 'dataStatistics表', 'POST'),
	(100, '2024-01-22 15:44:57.771', '2024-01-22 15:44:57.771', NULL, '/dataStatistics/deleteDataStatistics', '删除dataStatistics表', 'dataStatistics表', 'DELETE'),
	(101, '2024-01-22 15:44:57.989', '2024-01-22 15:44:57.989', NULL, '/dataStatistics/deleteDataStatisticsByIds', '批量删除dataStatistics表', 'dataStatistics表', 'DELETE'),
	(102, '2024-01-22 15:44:58.398', '2024-01-22 15:44:58.398', NULL, '/dataStatistics/updateDataStatistics', '更新dataStatistics表', 'dataStatistics表', 'PUT'),
	(103, '2024-01-22 15:44:58.675', '2024-01-22 15:44:58.675', NULL, '/dataStatistics/findDataStatistics', '根据ID获取dataStatistics表', 'dataStatistics表', 'GET'),
	(104, '2024-01-22 15:44:58.947', '2024-01-22 15:44:58.947', NULL, '/dataStatistics/getDataStatisticsList', '获取dataStatistics表列表', 'dataStatistics表', 'GET'),
	(105, '2024-02-01 15:16:35.635', '2024-02-01 15:16:35.635', NULL, '/userExt/createUserExt', '新增userExt表', 'userExt表', 'POST'),
	(106, '2024-02-01 15:16:35.822', '2024-02-01 15:16:35.822', NULL, '/userExt/deleteUserExt', '删除userExt表', 'userExt表', 'DELETE'),
	(107, '2024-02-01 15:16:36.057', '2024-02-01 15:16:36.057', NULL, '/userExt/deleteUserExtByIds', '批量删除userExt表', 'userExt表', 'DELETE'),
	(108, '2024-02-01 15:16:36.262', '2024-02-01 15:16:36.262', NULL, '/userExt/updateUserExt', '更新userExt表', 'userExt表', 'PUT'),
	(109, '2024-02-01 15:16:36.519', '2024-02-01 15:16:36.519', NULL, '/userExt/findUserExt', '根据ID获取userExt表', 'userExt表', 'GET'),
	(110, '2024-02-01 15:16:36.773', '2024-02-01 15:16:36.773', NULL, '/userExt/getUserExtList', '获取userExt表列表', 'userExt表', 'GET'),
	(111, '2024-02-01 16:41:02.813', '2024-02-01 16:41:58.965', NULL, '/userExt/v2/getList', '自建用户管理', 'userExtV2', 'GET'),
	(112, '2024-02-01 16:43:10.003', '2024-02-01 16:43:10.003', NULL, '/userExt/v2/create', '自建用户管理', 'userExtV2', 'POST'),
	(113, '2024-02-02 16:03:36.592', '2024-02-02 16:03:49.095', NULL, '/userExt/v2/update', '自建用户管理', 'userExtV2', 'PUT'),
	(114, '2024-02-03 09:17:21.594', '2024-02-03 09:17:21.594', NULL, '/user/batchQuery', '用户批量查询', '系统用户', 'POST'),
	(115, '2024-02-03 10:21:47.484', '2024-02-03 10:21:47.484', NULL, '/userDataUsage/createUserDataUsage', '新增userDataUsage表', 'userDataUsage表', 'POST'),
	(116, '2024-02-03 10:21:47.762', '2024-02-03 10:21:47.762', NULL, '/userDataUsage/deleteUserDataUsage', '删除userDataUsage表', 'userDataUsage表', 'DELETE'),
	(117, '2024-02-03 10:21:48.041', '2024-02-03 10:21:48.041', NULL, '/userDataUsage/deleteUserDataUsageByIds', '批量删除userDataUsage表', 'userDataUsage表', 'DELETE'),
	(118, '2024-02-03 10:21:48.393', '2024-02-03 10:21:48.393', NULL, '/userDataUsage/updateUserDataUsage', '更新userDataUsage表', 'userDataUsage表', 'PUT'),
	(119, '2024-02-03 10:21:48.625', '2024-02-03 10:21:48.625', NULL, '/userDataUsage/findUserDataUsage', '根据ID获取userDataUsage表', 'userDataUsage表', 'GET'),
	(120, '2024-02-03 10:21:48.871', '2024-02-03 10:21:48.871', NULL, '/userDataUsage/getUserDataUsageList', '获取userDataUsage表列表', 'userDataUsage表', 'GET'),
	(121, '2024-02-03 17:31:02.317', '2024-02-03 17:31:02.317', NULL, '/userExt/v2/dashboard', '用户面板', 'userExtV2', 'GET'),
	(122, '2024-02-04 09:49:00.972', '2024-02-04 09:49:00.972', NULL, '/announcement/createAnnouncement', '新增announcement表', 'announcement表', 'POST'),
	(123, '2024-02-04 09:49:01.173', '2024-02-04 09:49:01.173', NULL, '/announcement/deleteAnnouncement', '删除announcement表', 'announcement表', 'DELETE'),
	(124, '2024-02-04 09:49:01.406', '2024-02-04 09:49:01.406', NULL, '/announcement/deleteAnnouncementByIds', '批量删除announcement表', 'announcement表', 'DELETE'),
	(125, '2024-02-04 09:49:01.563', '2024-02-04 09:49:01.563', NULL, '/announcement/updateAnnouncement', '更新announcement表', 'announcement表', 'PUT'),
	(126, '2024-02-04 09:49:01.733', '2024-02-04 09:49:01.733', NULL, '/announcement/findAnnouncement', '根据ID获取announcement表', 'announcement表', 'GET'),
	(127, '2024-02-04 09:49:01.922', '2024-02-04 09:49:01.922', NULL, '/announcement/getAnnouncementList', '获取announcement表列表', 'announcement表', 'GET'),
	(128, '2024-02-28 14:58:50.177', '2024-02-28 14:58:50.177', NULL, '/v2rayNode/createV2rayNode', '新增v2rayNode表', 'v2rayNode表', 'POST'),
	(129, '2024-02-28 14:58:50.492', '2024-02-28 14:58:50.492', NULL, '/v2rayNode/deleteV2rayNode', '删除v2rayNode表', 'v2rayNode表', 'DELETE'),
	(130, '2024-02-28 14:58:50.731', '2024-02-28 14:58:50.731', NULL, '/v2rayNode/deleteV2rayNodeByIds', '批量删除v2rayNode表', 'v2rayNode表', 'DELETE'),
	(131, '2024-02-28 14:58:51.072', '2024-02-28 14:58:51.072', NULL, '/v2rayNode/updateV2rayNode', '更新v2rayNode表', 'v2rayNode表', 'PUT'),
	(132, '2024-02-28 14:58:51.496', '2024-02-28 14:58:51.496', NULL, '/v2rayNode/findV2rayNode', '根据ID获取v2rayNode表', 'v2rayNode表', 'GET'),
	(133, '2024-02-28 14:58:51.694', '2024-02-28 14:58:51.694', NULL, '/v2rayNode/getV2rayNodeList', '获取v2rayNode表列表', 'v2rayNode表', 'GET'),
	(134, '2024-03-11 14:42:58.094', '2024-03-11 14:42:58.094', NULL, '/v2rayNode/pushData', '主动推数据到节点', 'v2rayNode表', 'POST'),
	(136, '2024-03-14 10:35:15.032', '2024-03-14 10:35:15.032', NULL, '/inviteCode/createInviteCode', '新增inviteCode表', 'inviteCode表', 'POST'),
	(137, '2024-03-14 10:35:15.206', '2024-03-14 10:35:15.206', NULL, '/inviteCode/deleteInviteCode', '删除inviteCode表', 'inviteCode表', 'DELETE'),
	(138, '2024-03-14 10:35:15.427', '2024-03-14 10:35:15.427', NULL, '/inviteCode/deleteInviteCodeByIds', '批量删除inviteCode表', 'inviteCode表', 'DELETE'),
	(139, '2024-03-14 10:35:15.607', '2024-03-14 10:35:15.607', NULL, '/inviteCode/updateInviteCode', '更新inviteCode表', 'inviteCode表', 'PUT'),
	(140, '2024-03-14 10:35:15.774', '2024-03-14 10:35:15.774', NULL, '/inviteCode/findInviteCode', '根据ID获取inviteCode表', 'inviteCode表', 'GET'),
	(141, '2024-03-14 10:35:15.975', '2024-03-14 10:35:15.975', NULL, '/inviteCode/getInviteCodeList', '获取inviteCode表列表', 'inviteCode表', 'GET'),
	(142, '2024-03-14 11:58:27.638', '2024-03-14 11:58:27.638', NULL, '/inviteCode/getMyInviteCodeList', '获取用户所属的邀请码', 'inviteCode表', 'GET'),
	(143, '2024-03-14 13:45:55.849', '2024-03-14 13:45:55.849', NULL, '/inviteCode/deleteMyInviteCodeByIds', '批量删除我的邀请码', 'inviteCode表', 'DELETE');


CREATE TABLE IF NOT EXISTS `sys_authorities` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `authority_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`),
  UNIQUE KEY `authority_id` (`authority_id`)
) ENGINE=InnoDB AUTO_INCREMENT=95282 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `sys_authorities`;
INSERT INTO `sys_authorities` (`created_at`, `updated_at`, `deleted_at`, `authority_id`, `authority_name`, `parent_id`, `default_router`) VALUES
	('2024-03-07 11:36:39.076', '2024-03-14 16:51:54.758', NULL, 887, '开发模式', 888, 'dashboard'),
	('2024-01-18 16:24:21.351', '2024-03-14 16:51:48.056', NULL, 888, '管理员', 0, 'dashboard'),
	('2024-01-18 16:24:21.351', '2024-03-14 16:52:06.846', NULL, 9528, '用户', 0, 'dashboard');


CREATE TABLE IF NOT EXISTS `sys_authority_btns` (
  `authority_id` bigint unsigned DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint unsigned DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `sys_authority_btns`;


CREATE TABLE IF NOT EXISTS `sys_authority_menus` (
  `sys_base_menu_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `sys_authority_menus`;
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES
	(3, 887),
	(3, 888),
	(4, 887),
	(4, 888),
	(5, 887),
	(5, 888),
	(6, 887),
	(6, 888),
	(7, 887),
	(7, 888),
	(8, 887),
	(8, 888),
	(9, 887),
	(9, 888),
	(10, 887),
	(10, 888),
	(10, 9528),
	(11, 887),
	(12, 887),
	(13, 887),
	(14, 887),
	(15, 887),
	(15, 888),
	(16, 887),
	(17, 887),
	(18, 887),
	(18, 888),
	(19, 887),
	(20, 887),
	(21, 887),
	(23, 887),
	(23, 888),
	(24, 887),
	(25, 887),
	(26, 887),
	(27, 887),
	(28, 887),
	(29, 887),
	(30, 887),
	(31, 887),
	(31, 888),
	(31, 9528),
	(32, 887),
	(32, 888),
	(32, 9528),
	(33, 887),
	(33, 888),
	(33, 9528),
	(34, 887),
	(34, 888),
	(34, 9528),
	(35, 887),
	(35, 888),
	(36, 887),
	(36, 888),
	(37, 887),
	(37, 888),
	(38, 887),
	(38, 888),
	(39, 887),
	(39, 888),
	(39, 9528),
	(40, 887),
	(40, 888),
	(41, 887),
	(41, 888),
	(42, 887),
	(42, 888),
	(42, 9528),
	(43, 887),
	(43, 888),
	(44, 887),
	(44, 888),
	(44, 9528),
	(45, 887),
	(45, 888),
	(45, 9528),
	(46, 887),
	(46, 888),
	(46, 9528),
	(47, 887),
	(47, 888),
	(48, 887),
	(48, 888),
	(48, 9528);


CREATE TABLE IF NOT EXISTS `sys_auto_codes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `package_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '包名',
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '展示名',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_codes_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `sys_auto_codes`;
INSERT INTO `sys_auto_codes` (`id`, `created_at`, `updated_at`, `deleted_at`, `package_name`, `label`, `desc`) VALUES
	(1, '2024-01-22 15:43:47.525', '2024-01-22 15:43:47.525', NULL, 'datastatisticsmodel', '', ''),
	(2, '2024-02-01 15:15:43.821', '2024-02-01 15:15:43.821', NULL, 'userext', '', ''),
	(3, '2024-02-03 10:19:36.977', '2024-02-03 10:19:36.977', '2024-02-03 10:21:08.880', 'userDataUsage', '用户用量记录', ''),
	(4, '2024-02-03 10:21:28.103', '2024-02-03 10:21:28.103', NULL, 'userDataUsageModel', '', ''),
	(5, '2024-02-04 09:48:26.021', '2024-02-04 09:48:26.021', NULL, 'announcementModel', '', ''),
	(6, '2024-02-28 14:58:16.450', '2024-02-28 14:58:16.450', NULL, 'v2raynode', '', '');


CREATE TABLE IF NOT EXISTS `sys_auto_code_histories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `package` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `business_db` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `request_meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `auto_code_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `injection_meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `struct_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `struct_cn_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `api_ids` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `flag` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_code_histories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `sys_auto_code_histories`;
INSERT INTO `sys_auto_code_histories` (`id`, `created_at`, `updated_at`, `deleted_at`, `package`, `business_db`, `table_name`, `request_meta`, `auto_code_path`, `injection_meta`, `struct_name`, `struct_cn_name`, `api_ids`, `flag`) VALUES
	(1, '2024-01-22 15:44:59.320', '2024-01-22 15:44:59.320', NULL, 'datastatisticsmodel', '', 'data_statistics', '{"structName":"DataStatistics","tableName":"data_statistics","packageName":"dataStatistics","humpPackageName":"data_statistics","abbreviation":"dataStatistics","description":"dataStatistics表","autoCreateApiToSql":true,"autoCreateResource":false,"autoMoveFile":true,"businessDB":"","fields":[{"fieldName":"UserUuid","fieldDesc":"userUuid字段","fieldType":"string","fieldJson":"userUuid","dataTypeLong":"191","comment":"","columnName":"user_uuid","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false},{"fieldName":"Month","fieldDesc":"month字段","fieldType":"string","fieldJson":"month","dataTypeLong":"32","comment":"","columnName":"month","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false},{"fieldName":"Value","fieldDesc":"value字段","fieldType":"int","fieldJson":"value","dataTypeLong":"10","comment":"","columnName":"value","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false}],"HasTimer":false,"package":"datastatisticsmodel"}', '/home/wuguozhang/my-code/ladder-admin-v2/server/api/v1/datastatisticsmodel/data_statistics.go;/home/wuguozhang/my-code/ladder-admin-v2/server/model/datastatisticsmodel/data_statistics.go;/home/wuguozhang/my-code/ladder-admin-v2/server/model/datastatisticsmodel/request/data_statistics.go;/home/wuguozhang/my-code/ladder-admin-v2/server/router/datastatisticsmodel/data_statistics.go;/home/wuguozhang/my-code/ladder-admin-v2/server/service/datastatisticsmodel/data_statistics.go;/home/wuguozhang/my-code/ladder-admin-v2/web/src/api/dataStatistics.js;/home/wuguozhang/my-code/ladder-admin-v2/web/src/view/dataStatistics/dataStatisticsForm.vue;/home/wuguozhang/my-code/ladder-admin-v2/web/src/view/dataStatistics/dataStatistics.vue;', '/home/wuguozhang/my-code/ladder-admin-v2/server/api/v1/datastatisticsmodel/enter.go@ApiGroup@DataStatisticsApi;/home/wuguozhang/my-code/ladder-admin-v2/server/router/datastatisticsmodel/enter.go@RouterGroup@DataStatisticsRouter;/home/wuguozhang/my-code/ladder-admin-v2/server/service/datastatisticsmodel/enter.go@ServiceGroup@DataStatisticsService;', 'DataStatistics', 'dataStatistics表', '99;100;101;102;103;104;', 0),
	(2, '2024-02-01 15:16:37.174', '2024-02-01 15:16:37.174', NULL, 'userext', '', 'la_user_ext', '{"structName":"UserExt","tableName":"la_user_ext","packageName":"userExt","humpPackageName":"user_ext","abbreviation":"userExt","description":"userExt表","autoCreateApiToSql":true,"autoCreateResource":false,"autoMoveFile":true,"businessDB":"","fields":[{"fieldName":"Uuid","fieldDesc":"uuid字段","fieldType":"string","fieldJson":"uuid","dataTypeLong":"255","comment":"","columnName":"uuid","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false},{"fieldName":"ValidStart","fieldDesc":"validStart字段","fieldType":"time.Time","fieldJson":"validStart","dataTypeLong":"","comment":"","columnName":"valid_start","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false},{"fieldName":"ValidEnd","fieldDesc":"validEnd字段","fieldType":"time.Time","fieldJson":"validEnd","dataTypeLong":"","comment":"","columnName":"valid_end","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false}],"HasTimer":true,"package":"userext"}', '/home/wuguozhang/my-code/ladder-admin-v2/server/api/v1/userext/user_ext.go;/home/wuguozhang/my-code/ladder-admin-v2/server/model/userext/user_ext.go;/home/wuguozhang/my-code/ladder-admin-v2/server/model/userext/request/user_ext.go;/home/wuguozhang/my-code/ladder-admin-v2/server/router/userext/user_ext.go;/home/wuguozhang/my-code/ladder-admin-v2/server/service/userext/user_ext.go;/home/wuguozhang/my-code/ladder-admin-v2/web/src/api/userExt.js;/home/wuguozhang/my-code/ladder-admin-v2/web/src/view/userExt/userExtForm.vue;/home/wuguozhang/my-code/ladder-admin-v2/web/src/view/userExt/userExt.vue;', '/home/wuguozhang/my-code/ladder-admin-v2/server/api/v1/userext/enter.go@ApiGroup@UserExtApi;/home/wuguozhang/my-code/ladder-admin-v2/server/router/userext/enter.go@RouterGroup@UserExtRouter;/home/wuguozhang/my-code/ladder-admin-v2/server/service/userext/enter.go@ServiceGroup@UserExtService;', 'UserExt', 'userExt表', '105;106;107;108;109;110;', 0),
	(3, '2024-02-03 10:21:49.363', '2024-02-03 10:21:49.363', NULL, 'userDataUsageModel', '', 'la_user_data_usage', '{"structName":"UserDataUsage","tableName":"la_user_data_usage","packageName":"userDataUsage","humpPackageName":"user_data_usage","abbreviation":"userDataUsage","description":"userDataUsage表","autoCreateApiToSql":true,"autoCreateResource":false,"autoMoveFile":true,"businessDB":"","fields":[{"fieldName":"UserUuid","fieldDesc":"userUuid字段","fieldType":"string","fieldJson":"userUuid","dataTypeLong":"255","comment":"","columnName":"user_uuid","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false},{"fieldName":"StartDate","fieldDesc":"startDate字段","fieldType":"time.Time","fieldJson":"startDate","dataTypeLong":"","comment":"","columnName":"start_date","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false},{"fieldName":"EndDate","fieldDesc":"endDate字段","fieldType":"time.Time","fieldJson":"endDate","dataTypeLong":"","comment":"","columnName":"end_date","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false},{"fieldName":"StandardData","fieldDesc":"standardData字段","fieldType":"int","fieldJson":"standardData","dataTypeLong":"10","comment":"","columnName":"standard_data","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false},{"fieldName":"Usage","fieldDesc":"usage字段","fieldType":"int","fieldJson":"usage","dataTypeLong":"20","comment":"","columnName":"usage","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false}],"HasTimer":true,"package":"userDataUsageModel"}', '/home/wuguozhang/my-code/ladder-admin-v2/server/api/v1/userDataUsageModel/user_data_usage.go;/home/wuguozhang/my-code/ladder-admin-v2/server/model/userDataUsageModel/user_data_usage.go;/home/wuguozhang/my-code/ladder-admin-v2/server/model/userDataUsageModel/request/user_data_usage.go;/home/wuguozhang/my-code/ladder-admin-v2/server/router/userDataUsageModel/user_data_usage.go;/home/wuguozhang/my-code/ladder-admin-v2/server/service/userDataUsageModel/user_data_usage.go;/home/wuguozhang/my-code/ladder-admin-v2/web/src/api/userDataUsage.js;/home/wuguozhang/my-code/ladder-admin-v2/web/src/view/userDataUsage/userDataUsageForm.vue;/home/wuguozhang/my-code/ladder-admin-v2/web/src/view/userDataUsage/userDataUsage.vue;', '/home/wuguozhang/my-code/ladder-admin-v2/server/api/v1/userDataUsageModel/enter.go@ApiGroup@UserDataUsageApi;/home/wuguozhang/my-code/ladder-admin-v2/server/router/userDataUsageModel/enter.go@RouterGroup@UserDataUsageRouter;/home/wuguozhang/my-code/ladder-admin-v2/server/service/userDataUsageModel/enter.go@ServiceGroup@UserDataUsageService;', 'UserDataUsage', 'userDataUsage表', '115;116;117;118;119;120;', 0),
	(4, '2024-02-04 09:49:02.235', '2024-02-04 09:49:02.235', NULL, 'announcementModel', '', 'announcement', '{"structName":"Announcement","tableName":"announcement","packageName":"announcement","humpPackageName":"announcement","abbreviation":"announcement","description":"announcement表","autoCreateApiToSql":true,"autoCreateResource":false,"autoMoveFile":true,"businessDB":"","fields":[{"fieldName":"Title","fieldDesc":"title字段","fieldType":"string","fieldJson":"title","dataTypeLong":"255","comment":"","columnName":"title","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false},{"fieldName":"Content","fieldDesc":"content字段","fieldType":"string","fieldJson":"content","dataTypeLong":"","comment":"","columnName":"content","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false}],"HasTimer":false,"package":"announcementModel"}', '/home/wuguozhang/my-code/ladder-admin-v2/server/api/v1/announcementModel/announcement.go;/home/wuguozhang/my-code/ladder-admin-v2/server/model/announcementModel/announcement.go;/home/wuguozhang/my-code/ladder-admin-v2/server/model/announcementModel/request/announcement.go;/home/wuguozhang/my-code/ladder-admin-v2/server/router/announcementModel/announcement.go;/home/wuguozhang/my-code/ladder-admin-v2/server/service/announcementModel/announcement.go;/home/wuguozhang/my-code/ladder-admin-v2/web/src/api/announcement.js;/home/wuguozhang/my-code/ladder-admin-v2/web/src/view/announcement/announcementForm.vue;/home/wuguozhang/my-code/ladder-admin-v2/web/src/view/announcement/announcement.vue;', '/home/wuguozhang/my-code/ladder-admin-v2/server/api/v1/announcementModel/enter.go@ApiGroup@AnnouncementApi;/home/wuguozhang/my-code/ladder-admin-v2/server/router/announcementModel/enter.go@RouterGroup@AnnouncementRouter;/home/wuguozhang/my-code/ladder-admin-v2/server/service/announcementModel/enter.go@ServiceGroup@AnnouncementService;', 'Announcement', 'announcement表', '122;123;124;125;126;127;', 0),
	(5, '2024-02-28 14:58:52.042', '2024-02-28 14:58:52.042', NULL, 'v2raynode', '', 'la_v2ray_node', '{"structName":"V2rayNode","tableName":"la_v2ray_node","packageName":"v2rayNode","humpPackageName":"v2ray_node","abbreviation":"v2rayNode","description":"v2rayNode表","autoCreateApiToSql":true,"autoCreateResource":false,"autoMoveFile":true,"businessDB":"","fields":[{"fieldName":"UniqueId","fieldDesc":"uniqueId字段","fieldType":"string","fieldJson":"uniqueId","dataTypeLong":"64","comment":"","columnName":"unique_id","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false},{"fieldName":"ConfigRaw","fieldDesc":"configRaw字段","fieldType":"string","fieldJson":"configRaw","dataTypeLong":"","comment":"","columnName":"config_raw","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false},{"fieldName":"Name","fieldDesc":"name字段","fieldType":"string","fieldJson":"name","dataTypeLong":"64","comment":"","columnName":"name","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false},{"fieldName":"Secret","fieldDesc":"secret字段","fieldType":"string","fieldJson":"secret","dataTypeLong":"64","comment":"","columnName":"secret","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false},{"fieldName":"SecretIv","fieldDesc":"secretIv字段","fieldType":"string","fieldJson":"secretIv","dataTypeLong":"64","comment":"","columnName":"secret_iv","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false},{"fieldName":"VmessPort","fieldDesc":"vmessPort字段","fieldType":"int","fieldJson":"vmessPort","dataTypeLong":"","comment":"","columnName":"vmess_port","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false},{"fieldName":"RpcPort","fieldDesc":"rpcPort字段","fieldType":"int","fieldJson":"rpcPort","dataTypeLong":"","comment":"","columnName":"rpc_port","fieldSearchType":"","dictType":"","require":false,"errorText":"","clearable":true,"sort":false}],"HasTimer":false,"package":"v2raynode"}', '/home/wuguozhang/my-code/ladder-admin-v2/server/api/v1/v2raynode/v2ray_node.go;/home/wuguozhang/my-code/ladder-admin-v2/server/model/v2raynode/v2ray_node.go;/home/wuguozhang/my-code/ladder-admin-v2/server/model/v2raynode/request/v2ray_node.go;/home/wuguozhang/my-code/ladder-admin-v2/server/router/v2raynode/v2ray_node.go;/home/wuguozhang/my-code/ladder-admin-v2/server/service/v2raynode/v2ray_node.go;/home/wuguozhang/my-code/ladder-admin-v2/web/src/api/v2rayNode.js;/home/wuguozhang/my-code/ladder-admin-v2/web/src/view/v2rayNode/v2rayNodeForm.vue;/home/wuguozhang/my-code/ladder-admin-v2/web/src/view/v2rayNode/v2rayNode.vue;', '/home/wuguozhang/my-code/ladder-admin-v2/server/api/v1/v2raynode/enter.go@ApiGroup@V2rayNodeApi;/home/wuguozhang/my-code/ladder-admin-v2/server/router/v2raynode/enter.go@RouterGroup@V2rayNodeRouter;/home/wuguozhang/my-code/ladder-admin-v2/server/service/v2raynode/enter.go@ServiceGroup@V2rayNodeService;', 'V2rayNode', 'v2rayNode表', '128;129;130;131;132;133;', 0);


CREATE TABLE IF NOT EXISTS `sys_base_menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` bigint unsigned DEFAULT NULL,
  `parent_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `sys_base_menus`;
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `active_name`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES
	(2, '2024-01-18 16:24:26.712', '2024-03-07 11:35:32.910', NULL, 0, '0', 'about', 'about', 1, 'view/about/index.vue', 9, '', 0, 0, '关于我们', 'info-filled', 0),
	(3, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '0', 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, '', 0, 0, '超级管理员', 'user', 0),
	(4, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '3', 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, '', 0, 0, '角色管理', 'avatar', 0),
	(5, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '3', 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, '', 1, 0, '菜单管理', 'tickets', 0),
	(6, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '3', 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, '', 1, 0, 'api管理', 'platform', 0),
	(7, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '3', 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, '', 0, 0, '用户管理', 'coordinate', 0),
	(8, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '3', 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, '', 0, 0, '字典管理', 'notebook', 0),
	(9, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '3', 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, '', 0, 0, '操作历史', 'pie-chart', 0),
	(10, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '0', 'person', 'person', 1, 'view/person/person.vue', 4, '', 0, 0, '个人信息', 'message', 0),
	(11, '2024-01-18 16:24:26.712', '2024-03-07 11:38:23.519', NULL, 0, '0', 'example', 'example', 0, 'view/example/index.vue', 7, '', 0, 0, '示例文件', 'management', 0),
	(12, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '11', 'upload', 'upload', 0, 'view/example/upload/upload.vue', 5, '', 0, 0, '媒体库（上传下载）', 'upload', 0),
	(13, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '11', 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, '', 0, 0, '断点续传', 'upload-filled', 0),
	(14, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '11', 'customer', 'customer', 0, 'view/example/customer/customer.vue', 7, '', 0, 0, '客户列表（资源示例）', 'avatar', 0),
	(15, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '0', 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, '', 0, 0, '系统工具', 'tools', 0),
	(16, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '15', 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, '', 1, 0, '代码生成器', 'cpu', 0),
	(17, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '15', 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 2, '', 1, 0, '表单生成器', 'magic-stick', 0),
	(18, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '15', 'system', 'system', 0, 'view/systemTools/system/system.vue', 3, '', 0, 0, '系统配置', 'operation', 0),
	(19, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '15', 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 1, '', 0, 0, '自动化代码管理', 'magic-stick', 0),
	(20, '2024-01-18 16:24:26.712', '2024-03-07 11:38:09.563', NULL, 0, '15', 'autoCodeEdit/:id', 'autoCodeEdit', 0, 'view/systemTools/autoCode/index.vue', 0, '', 0, 0, '自动化代码-${id}', 'magic-stick', 0),
	(21, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '15', 'autoPkg', 'autoPkg', 0, 'view/systemTools/autoPkg/autoPkg.vue', 0, '', 0, 0, '自动化package', 'folder', 0),
	(23, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '0', 'state', 'state', 0, 'view/system/state.vue', 8, '', 0, 0, '服务器状态', 'cloudy', 0),
	(24, '2024-01-18 16:24:26.712', '2024-03-07 11:38:17.250', NULL, 0, '0', 'plugin', 'plugin', 0, 'view/routerHolder.vue', 6, '', 0, 0, '插件系统', 'cherry', 0),
	(25, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '24', 'https://plugin.gin-vue-admin.com/', 'https://plugin.gin-vue-admin.com/', 0, 'https://plugin.gin-vue-admin.com/', 0, '', 0, 0, '插件市场', 'shop', 0),
	(26, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '24', 'installPlugin', 'installPlugin', 0, 'view/systemTools/installPlugin/index.vue', 1, '', 0, 0, '插件安装', 'box', 0),
	(27, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '24', 'autoPlug', 'autoPlug', 0, 'view/systemTools/autoPlug/autoPlug.vue', 2, '', 0, 0, '插件模板', 'folder', 0),
	(28, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '24', 'pubPlug', 'pubPlug', 0, 'view/systemTools/pubPlug/pubPlug.vue', 3, '', 0, 0, '打包插件', 'files', 0),
	(29, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '24', 'plugin-email', 'plugin-email', 0, 'plugin/email/view/index.vue', 4, '', 0, 0, '邮件插件', 'message', 0),
	(30, '2024-01-18 16:24:26.712', '2024-01-18 16:24:26.712', NULL, 0, '15', 'exportTemplate', 'exportTemplate', 0, 'view/systemTools/exportTemplate/exportTemplate.vue', 10, '', 0, 0, '表格模板', 'reading', 0),
	(31, '2024-01-18 16:45:51.241', '2024-01-22 14:35:44.487', NULL, 0, '0', 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 0, '', 0, 0, '首页', 'flag', 0),
	(32, '2024-01-22 14:40:17.390', '2024-01-22 15:00:25.179', NULL, 0, '0', 'my-root', 'my-root', 0, 'view/routerHolder.vue', 1, '', 0, 0, '我的', 'user', 0),
	(33, '2024-01-22 14:42:55.468', '2024-02-03 15:28:46.931', NULL, 0, '32', 'my-history', 'my-history', 0, 'view/myHistory/index.vue', 30, '', 0, 0, '历史流量数据', '', 0),
	(34, '2024-01-22 14:47:59.354', '2024-02-03 15:28:16.425', NULL, 0, '32', 'my', 'my', 0, 'view/my/my.vue', 10, '', 0, 0, '基础信息', '', 0),
	(35, '2024-01-22 15:48:33.774', '2024-02-04 09:52:01.112', NULL, 0, '36', 'data-stat', 'data-stat', 0, 'view/dataStatistics/index.vue', 30, '', 0, 0, '数据列表', '', 0),
	(36, '2024-01-31 17:31:41.672', '2024-01-31 17:31:41.672', NULL, 0, '0', 'sysadmin', 'sysadmin', 0, 'view/routerHolder.vue', 2, '', 0, 0, '系统管理', 'setting', 0),
	(37, '2024-01-31 17:34:07.032', '2024-02-04 09:51:21.281', NULL, 0, '36', 'useradmin', 'useradmin', 0, 'view/sysadmin/user/index.vue', 10, '', 0, 0, '用户管理', '', 0),
	(38, '2024-02-03 10:38:43.143', '2024-02-04 09:51:44.269', NULL, 0, '36', 'usage', 'usage', 0, 'view/userDataUsage/userDataUsage.vue', 20, '', 0, 0, '用量列表', '', 0),
	(39, '2024-02-03 15:27:27.070', '2024-02-03 15:28:37.247', NULL, 0, '32', 'my-usage', 'my-usage', 0, 'view/myUsage/index.vue', 20, '', 0, 0, '我的用量数据', '', 0),
	(40, '2024-02-04 09:52:33.230', '2024-02-04 09:52:33.230', NULL, 0, '36', 'announcement', 'announcement', 0, 'view/announcement/announcement.vue', 40, '', 0, 0, '公告管理', '', 0),
	(41, '2024-02-04 10:23:39.494', '2024-02-04 14:03:02.249', NULL, 0, '36', 'announcement/edit', 'announcement-edit', 1, 'view/announcement/edit.vue', 0, '', 0, 0, '编辑公告', '', 0),
	(42, '2024-02-04 15:30:29.702', '2024-03-08 16:08:05.507', NULL, 0, '0', 'guide', 'guide', 1, 'view/routerHolder.vue', 2, '', 0, 0, '教程', 'guide', 0),
	(43, '2024-02-28 15:05:56.759', '2024-02-28 15:06:11.961', NULL, 0, '36', 'v2ray-node', 'v2ray-node', 0, 'view/v2rayNode/v2rayNode.vue', 50, '', 0, 0, '节点管理', '', 0),
	(44, '2024-03-07 11:43:49.876', '2024-03-07 11:47:32.720', NULL, 0, '42', 'how', 'how', 0, 'view/guide/index.vue', 1, '', 0, 0, '教程全文', 'guide', 0),
	(45, '2024-03-07 11:44:25.397', '2024-03-07 11:45:41.260', NULL, 0, '42', 'download', 'download', 0, 'view/guide/download.vue', 10, '', 0, 0, '下载', 'download', 0),
	(46, '2024-03-14 10:47:16.403', '2024-03-14 10:47:16.403', NULL, 0, '0', 'invite-code', 'invite-code', 0, 'view/routerHolder.vue', 5, '', 0, 0, '邀请码', 'user', 0),
	(47, '2024-03-14 10:47:59.241', '2024-03-14 10:47:59.241', NULL, 0, '46', 'invite-code-list', 'invite-code-list', 0, 'view/inviteCode/inviteCode.vue', 10, '', 0, 0, '邀请码列表', '', 0),
	(48, '2024-03-14 11:50:34.762', '2024-03-14 11:50:34.762', NULL, 0, '46', 'my-invite-code', 'my-invite-code', 0, 'view/inviteCode/myInviteCode.vue', 20, '', 0, 0, '我的邀请码', '', 0);


CREATE TABLE IF NOT EXISTS `sys_base_menu_btns` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_btns_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `sys_base_menu_btns`;


CREATE TABLE IF NOT EXISTS `sys_base_menu_parameters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `sys_base_menu_parameters`;


CREATE TABLE IF NOT EXISTS `sys_data_authority_id` (
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `sys_data_authority_id`;
INSERT INTO `sys_data_authority_id` (`sys_authority_authority_id`, `data_authority_id_authority_id`) VALUES
	(888, 888),
	(888, 8881),
	(888, 9528),
	(9528, 8881),
	(9528, 9528);


CREATE TABLE IF NOT EXISTS `sys_dictionaries` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `sys_dictionaries`;
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES
	(1, '2024-01-18 16:24:22.985', '2024-01-18 16:24:23.405', NULL, '性别', 'gender', 1, '性别字典'),
	(2, '2024-01-18 16:24:22.985', '2024-01-18 16:24:23.970', NULL, '数据库int类型', 'int', 1, 'int类型对应的数据库类型'),
	(3, '2024-01-18 16:24:22.985', '2024-01-18 16:24:24.481', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型'),
	(4, '2024-01-18 16:24:22.985', '2024-01-18 16:24:24.983', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型'),
	(5, '2024-01-18 16:24:22.985', '2024-01-18 16:24:25.475', NULL, '数据库字符串', 'string', 1, '数据库字符串'),
	(6, '2024-01-18 16:24:22.985', '2024-01-18 16:24:26.042', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型');


CREATE TABLE IF NOT EXISTS `sys_dictionary_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '展示值',
  `value` bigint DEFAULT NULL COMMENT '字典值',
  `extend` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '扩展值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint unsigned DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `sys_dictionary_details`;
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `extend`, `status`, `sort`, `sys_dictionary_id`) VALUES
	(1, '2024-01-18 16:24:23.525', '2024-01-18 16:24:23.525', NULL, '男', 1, '', 1, 1, 1),
	(2, '2024-01-18 16:24:23.525', '2024-01-18 16:24:23.525', NULL, '女', 2, '', 1, 2, 1),
	(3, '2024-01-18 16:24:24.057', '2024-01-18 16:24:24.057', NULL, 'smallint', 1, 'mysql', 1, 1, 2),
	(4, '2024-01-18 16:24:24.057', '2024-01-18 16:24:24.057', NULL, 'mediumint', 2, 'mysql', 1, 2, 2),
	(5, '2024-01-18 16:24:24.057', '2024-01-18 16:24:24.057', NULL, 'int', 3, 'mysql', 1, 3, 2),
	(6, '2024-01-18 16:24:24.057', '2024-01-18 16:24:24.057', NULL, 'bigint', 4, 'mysql', 1, 4, 2),
	(7, '2024-01-18 16:24:24.057', '2024-01-18 16:24:24.057', NULL, 'int2', 5, 'pgsql', 1, 5, 2),
	(8, '2024-01-18 16:24:24.057', '2024-01-18 16:24:24.057', NULL, 'int4', 6, 'pgsql', 1, 6, 2),
	(9, '2024-01-18 16:24:24.057', '2024-01-18 16:24:24.057', NULL, 'int6', 7, 'pgsql', 1, 7, 2),
	(10, '2024-01-18 16:24:24.057', '2024-01-18 16:24:24.057', NULL, 'int8', 8, 'pgsql', 1, 8, 2),
	(11, '2024-01-18 16:24:24.564', '2024-01-18 16:24:24.564', NULL, 'date', 0, '', 1, 0, 3),
	(12, '2024-01-18 16:24:24.564', '2024-01-18 16:24:24.564', NULL, 'time', 1, 'mysql', 1, 1, 3),
	(13, '2024-01-18 16:24:24.564', '2024-01-18 16:24:24.564', NULL, 'year', 2, 'mysql', 1, 2, 3),
	(14, '2024-01-18 16:24:24.564', '2024-01-18 16:24:24.564', NULL, 'datetime', 3, 'mysql', 1, 3, 3),
	(15, '2024-01-18 16:24:24.564', '2024-01-18 16:24:24.564', NULL, 'timestamp', 5, 'mysql', 1, 5, 3),
	(16, '2024-01-18 16:24:24.564', '2024-01-18 16:24:24.564', NULL, 'timestamptz', 6, 'pgsql', 1, 5, 3),
	(17, '2024-01-18 16:24:25.082', '2024-01-18 16:24:25.082', NULL, 'float', 0, '', 1, 0, 4),
	(18, '2024-01-18 16:24:25.082', '2024-01-18 16:24:25.082', NULL, 'double', 1, 'mysql', 1, 1, 4),
	(19, '2024-01-18 16:24:25.082', '2024-01-18 16:24:25.082', NULL, 'decimal', 2, 'mysql', 1, 2, 4),
	(20, '2024-01-18 16:24:25.082', '2024-01-18 16:24:25.082', NULL, 'numeric', 3, 'pgsql', 1, 3, 4),
	(21, '2024-01-18 16:24:25.082', '2024-01-18 16:24:25.082', NULL, 'smallserial', 4, 'pgsql', 1, 4, 4),
	(22, '2024-01-18 16:24:25.578', '2024-01-18 16:24:25.578', NULL, 'char', 0, '', 1, 0, 5),
	(23, '2024-01-18 16:24:25.578', '2024-01-18 16:24:25.578', NULL, 'varchar', 1, 'mysql', 1, 1, 5),
	(24, '2024-01-18 16:24:25.578', '2024-01-18 16:24:25.578', NULL, 'tinyblob', 2, 'mysql', 1, 2, 5),
	(25, '2024-01-18 16:24:25.578', '2024-01-18 16:24:25.578', NULL, 'tinytext', 3, 'mysql', 1, 3, 5),
	(26, '2024-01-18 16:24:25.578', '2024-01-18 16:24:25.578', NULL, 'text', 4, 'mysql', 1, 4, 5),
	(27, '2024-01-18 16:24:25.578', '2024-01-18 16:24:25.578', NULL, 'blob', 5, 'mysql', 1, 5, 5),
	(28, '2024-01-18 16:24:25.578', '2024-01-18 16:24:25.578', NULL, 'mediumblob', 6, 'mysql', 1, 6, 5),
	(29, '2024-01-18 16:24:25.578', '2024-01-18 16:24:25.578', NULL, 'mediumtext', 7, 'mysql', 1, 7, 5),
	(30, '2024-01-18 16:24:25.578', '2024-01-18 16:24:25.578', NULL, 'longblob', 8, 'mysql', 1, 8, 5),
	(31, '2024-01-18 16:24:25.578', '2024-01-18 16:24:25.578', NULL, 'longtext', 9, 'mysql', 1, 9, 5),
	(32, '2024-01-18 16:24:26.146', '2024-01-18 16:24:26.146', NULL, 'tinyint', 1, 'mysql', 1, 0, 6),
	(33, '2024-01-18 16:24:26.146', '2024-01-18 16:24:26.146', NULL, 'bool', 2, 'pgsql', 1, 0, 6);

CREATE TABLE IF NOT EXISTS `sys_export_templates` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板名称',
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '表名称',
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板标识',
  `template_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_templates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DELETE FROM `sys_export_templates`;


CREATE TABLE IF NOT EXISTS `sys_operation_records` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求路径',
  `status` bigint DEFAULT NULL COMMENT '请求状态',
  `latency` bigint DEFAULT NULL COMMENT '延迟',
  `agent` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '代理',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '错误信息',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '请求Body',
  `resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '响应Body',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_operation_records_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=781 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `sys_operation_records`;


CREATE TABLE IF NOT EXISTS `sys_users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '系统用户' COMMENT '用户昵称',
  `side_mode` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'dark' COMMENT '用户侧边主题',
  `header_img` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户头像',
  `base_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '#fff' COMMENT '基础颜色',
  `active_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '#1890ff' COMMENT '活跃颜色',
  `authority_id` bigint unsigned DEFAULT '888' COMMENT '用户角色ID',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint DEFAULT '1' COMMENT '用户是否被冻结 1正常 2冻结',
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`),
  KEY `idx_sys_users_uuid` (`uuid`),
  KEY `idx_sys_users_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `sys_users`;
INSERT INTO `sys_users` (`id`, `created_at`, `updated_at`, `deleted_at`, `uuid`, `username`, `password`, `nick_name`, `side_mode`, `header_img`, `base_color`, `active_color`, `authority_id`, `phone`, `email`, `enable`) VALUES
	(1, '2024-01-18 16:24:27.211', '2024-03-07 11:07:14.066', NULL, '41b8f634-5efc-ce05-ed70-203bd2cc005c', 'admin', '$2a$10$6TbMPaOvNPu02AL.sJ3uP.Nch8PajOGMUsPs9ZeH0CJ1cP1ZxoZ4m', 'admin用户', 'dark', '', '#fff', '#1890ff', 888, '15611111111', 'admin@ladder-admin.org', 1);


CREATE TABLE IF NOT EXISTS `sys_user_authority` (
  `sys_user_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DELETE FROM `sys_user_authority`;
INSERT INTO `sys_user_authority` (`sys_user_id`, `sys_authority_authority_id`) VALUES
	(1, 888),
	(1, 9528);

