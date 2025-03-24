package initializers

// ---------------
// Docker
// ---------------

var DATCOM_DOCKER bool = false

// ---------------
// Token & Keys
// ---------------

var ENUM_DATCOM_ROLE_MEMBER uint = 5   // datcom member, maximum value for user.Role
var ENUM_DATCOM_COURSE_MEMBER uint = 2 // datcom course, maximum value for user.Course

var TOKEN_EXPIRES_HOURS float64 = 168 // expiration time for the token in hours
var SIZE_DOCUMENT_KEY int8 = 8        // document upload key size

// ---------------
// Database & Authentication
// ---------------

// WARNING: CAREFUL EDIT
var DATCOM_ADMIN_USER string = "admin"
var DATCOM_ADMIN_PWD string = "3b74658433fcf858ad4442cba0a984db"

// WARNING: DO NOT EDIT
var DATCOM_DB_PWD string = "0ac0bf6efe"
