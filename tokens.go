/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

// TokenSupply gets token supply
func (c *Client) TokenSupply(address string) (balance *BigInt, err error) {
	param := M{
		"contractaddress": address,
	}
	balance = new(BigInt)
	err = c.call("stats", "tokensupply", param, &balance)
	return
}

// TokenCSupply gets token circulating supply
func (c *Client) TokenCSupply(address string) (balance *BigInt, err error) {
	param := M{
		"contractaddress": address,
	}
	balance = new(BigInt)
	err = c.call("stats", "tokenCsupply", param, &balance)
	return
}
