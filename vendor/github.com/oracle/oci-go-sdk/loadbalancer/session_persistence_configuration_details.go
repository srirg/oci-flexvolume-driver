// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SessionPersistenceConfigurationDetails The configuration details for implementing session persistence. Session persistence enables the Load Balancing
// Service to direct any number of requests that originate from a single logical client to a single backend web server.
// For more information, see [Session Persistence]({{DOC_SERVER_URL}}/Content/Balance/Reference/sessionpersistence.htm).
// To disable session persistence on a running load balancer, use the
// UpdateBackendSet operation and specify "null" for the
// `SessionPersistenceConfigurationDetails` object.
// Example: `SessionPersistenceConfigurationDetails: null`
type SessionPersistenceConfigurationDetails struct {

	// The name of the cookie used to detect a session initiated by the backend server. Use '*' to specify
	// that any cookie set by the backend causes the session to persist.
	// Example: `myCookieName`
	CookieName *string `mandatory:"true" json:"cookieName"`

	// Whether the load balancer is prevented from directing traffic from a persistent session client to
	// a different backend server if the original server is unavailable. Defaults to false.
	// Example: `true`
	DisableFallback *bool `mandatory:"false" json:"disableFallback"`
}

func (m SessionPersistenceConfigurationDetails) String() string {
	return common.PointerString(m)
}
