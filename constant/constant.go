package constant

// All constant for pagination and order
const (
	// DefaultLimit is default limit for pagination
	DefaultLimit = 10

	// DefaultPage is default page for pagination
	DefaultPage = 1

	// DefaultSort is default sort for pagination
	DefaultSort = "id"

	// DefaultOrder is default order for pagination
	DefaultOrder = "asc"

	// DefaultCode is default code for response
	DefaultCode = 200
)

// All constant message for response
const (

	// DefaultMessage is default message for response
	DefaultMessage = "Success."

	// SuccessMessage is the message returned when the request is success
	SuccessMessage = "The request was processed successfully."

	// FailedMessage is the message returned when the request is failed
	FailedMessage = "The request failed to process."

	// DataFound is the message returned when data is found
	DataFound = "Data found."

	// DataNotFound is the message returned when data is not found
	DataNotFound = "Data not found."

	// CannotProcessRequest is the message returned when the request cannot be processed
	CannotProcessRequest = "Cannot process request."

	// InvalidRequest is the message returned when the request is invalid
	InvalidRequest = "Invalid request."

	// SuccessCreateData is the message returned when the data is successfully created
	SuccessCreateData = "Successfully created new data."

	// FailedCreateData is the message returned when the data failed to be created
	FailedCreateData = "Failed to create new data."

	// SuccessGetData is the message returned when the data is successfully retrieved
	SuccessGetData = "Successfully retrieved data."

	// SuccessUpdateData is the message returned when the data is successfully updated
	SuccessUpdateData = "Successfully updated data."

	// FailedUpdateData is the message returned when the data failed to be updated
	FailedUpdateData = "Failed to update data."

	// SuccessDeleteData is the message returned when the data is successfully deleted
	SuccessDeleteData = "Successfully deleted data."

	// FailedDeleteData is the message returned when the data failed to be deleted
	FailedDeleteData = "Failed to delete data."
)
