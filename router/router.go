/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/polarismesh/polaris-console/bootstrap"
	"github.com/polarismesh/polaris-console/handlers"
)

/**
 * @brief 路由请求
 */
func Router(config *bootstrap.Config) {
	r := gin.Default()
	// 加载静态资源
	r.Static("/static", config.WebServer.WebPath+"static")

	// 加载界面
	r.LoadHTMLGlob(config.WebServer.WebPath + "index.html")
	r.GET("/", handlers.PolarisPage(&config.OAAuthority))

	// 获取部门数据
	r.GET("/HRFoundation-Unit", handlers.GetDepartment(&config.HRData, &config.OAAuthority))
	// 通过企业微信名获取部门数据
	r.GET("/getStaffDept", handlers.ReverseProxyForDepartment(&config.HRData, &config.OAAuthority))
	// 查询路由/限流/熔断日志记录
	r.POST("/log/search/elasticsearch", handlers.ReverseProxyForLogRecord(&config.ZhiYan))

	// 后端server路由组
	v1 := r.Group(config.WebServer.RequestURL)
	// 创建命名空间
	v1.POST("/namespaces", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, false))
	// 创建服务
	v1.POST("/services", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, false))
	// 创建服务别名
	v1.POST("/service/alias", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, true))
	// 创建服务实例
	v1.POST("/instances", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, true))
	// 创建路由
	v1.POST("/routings", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, true))
	// 创建限流规则
	v1.POST("/ratelimits", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, true))
	// 创建熔断规则
	v1.POST("/circuitbreakers", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, false))
	// 创建熔断规则版本
	v1.POST("/circuitbreakers/version", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, true))
	// 发布熔断规则（已经在前端对负责人信息进行校验）
	v1.POST("/circuitbreakers/release", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, false))

	// 查看资源
	v1.GET("/:resource", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, false))
	// 查看服务绑定的熔断规则
	v1.GET("/:resource/circuitbreaker", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, false))
	// 查看master版本的规则
	v1.GET("/:resource/master", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, false))
	// 查看已发布的规则
	v1.GET("/:resource/release", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, false))
	// 查看规则的所有版本
	v1.GET("/:resource/versions", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, false))
	// 查看服务和实例个数
	v1.GET("/:resource/count", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, false))
	// 查看服务别名
	v1.GET("/:resource/aliases", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, false))
	// 查看Token，需要鉴权
	v1.GET("/:resource/token", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, true))

	// 修改资源
	v1.PUT("/:resource", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, true))
	// 修改Token
	v1.PUT("/:resource/token", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, true))

	// 删除命名空间
	v1.POST("/namespaces/delete", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, true))
	// 删除服务或服务别名
	v1.POST("/services/delete", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, true))
	// 删除服务实例
	v1.POST("/instances/delete", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, true))
	// 删除路由
	v1.POST("/routings/delete", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, true))
	// 删除限流规则
	v1.POST("ratelimits/delete", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, true))
	// 删除熔断规则
	v1.POST("circuitbreakers/delete", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, true))
	// 解绑熔断规则
	v1.POST("circuitbreakers/unbind", handlers.ReverseProxyForServer(&config.PolarisServer, &config.OAAuthority, true))

	// 监控请求路由组
	mv1 := r.Group(config.WebServer.MonitorURL)
	mv1.GET("/query_range", handlers.ReverseProxyForMonitorServer(&config.MonitorServer))

	address := fmt.Sprintf("%v:%v", config.WebServer.ListenIP, config.WebServer.ListenPort)
	r.Run(address)
}
