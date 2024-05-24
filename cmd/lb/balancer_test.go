package main

import (
  "fmt"
  . "gopkg.in/check.v1"
  "testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestScheme(c *C) {
  *https = true
  c.Assert(scheme(), Equals, "https")
  fmt.Printf("Scheme is %s when https is true\n", scheme())

  *https = false
  c.Assert(scheme(), Equals, "http")
  fmt.Printf("Scheme is %s when https is false\n", scheme())
}

func (s *MySuite) TestHash(c *C) {
  str := "testString"
  c.Assert(hash(str), FitsTypeOf, uint32(0))
  fmt.Printf("Hash of %s is %d\n", str, hash(str))
}

func (s *MySuite) TestLoadBalancer(c *C) {
  path1 := "/some/test/path1"
  path2 := "/some/test/path2"

  for i := range healthyServers {
    healthyServers[i] = true
    fmt.Println("healthyServers after iteration", i, ":", healthyServers)
  }

  firstServerPath1 := chooseServer(path1)
  c.Assert(firstServerPath1, Not(Equals), "")
  fmt.Printf("First server chosen for path1 is %s\n", firstServerPath1)

  firstServerPath2 := chooseServer(path2)
  c.Assert(firstServerPath2, Not(Equals), "")
  fmt.Printf("First server chosen for path2 is %s\n", firstServerPath2)

  for i := 0; i < 10; i++ {
    serverPath1 := chooseServer(path1)
    c.Assert(serverPath1, Equals, firstServerPath1)
    fmt.Printf("Server chosen on iteration %d for path1 is %s\n", i, serverPath1)

    serverPath2 := chooseServer(path2)
    c.Assert(serverPath2, Equals, firstServerPath2)
    fmt.Printf("Server chosen on iteration %d for path2 is %s\n", i, serverPath2)
  }
}

