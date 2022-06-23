/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package redis

import (
	"strconv"

	"github.com/go-redis/redis"
)

type Param struct {
	Address  string
	Password string
	DB       string
}

func (c *Param) New(impl *Redis) (*Redis, error) {
	if c.DB == "" {
		c.DB = "0"
	}
	dbInt, err := strconv.Atoi(c.DB)
	if err != nil {
		return impl, err
	}
	client := redis.NewClient(&redis.Options{
		Addr:     c.Address,
		Password: c.Password,
		DB:       dbInt,
	})
	_, err = client.Ping().Result()
	if err != nil {
		return impl, err
	}
	impl.client = client
	return impl, nil
}
