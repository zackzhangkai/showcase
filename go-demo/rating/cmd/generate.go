/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"gorm.io/gen"
	"ratings/biz/dal/mysql"
	genModel "ratings/biz/model/model"

	// reuse your gorm db
	// init db
	_ "ratings/biz/dal"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./biz/model/db",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	// reuse your gorm db
	g.UseDB(mysql.DB)

	// Generate struct `User` based on table `users`
	g.GenerateModel("ratings")

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(genModel.Rating{})

	// Generate the code
	g.Execute()
}
