package initializers

// ---------------
// Docker
// ---------------

var DATCOM_DOCKER bool = true

// ---------------
// Token & Keys
// ---------------

var ENUM_DATCOM_ROLE_MEMBER uint = 6  // datcom member, maximum value for user.Role
var TOKEN_EXPIRES_HOURS float64 = 168 // expiration time for the token in hours
var SIZE_DOCUMENT_KEY int8 = 8        // document upload key size

// ---------------
// Database & Authentication
// ---------------

// WARNING: CAREFUL EDIT
var DATCOM_ADMIN_USER string = "admin"
var DATCOM_ADMIN_PWD string = "c39a690203eb9300001addefb63461c5"

// WARNING: DO NOT EDIT
var DATCOM_DB_PWD string = "27ba0f8d32"
