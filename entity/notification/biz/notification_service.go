//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/godine/entity/notification/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// ListNotificationsOptions defines the options for listing notifications.
type ListNotificationsOptions struct {
	// Page specifies the page number for pagination.
	Page int `form:"page" default:"1" minimum:"1"`

	// Size specifies the number of items per page.
	Size int `form:"size" default:"10" minimum:"1" maximum:"100"`
}

// INotificationBiz defines the interface for notification management operations.
// This interface abstracts the notification management functionality, making it
// easier to manage notifications and their statuses. The design allows for flexibility,
// scalability, and maintainability of the notification operations in the GoDine system.
type INotificationBiz interface {
	// CreateNotification creates a new notification. This method is responsible for
	// initializing a new notification entity with provided details such as UserID,
	// OrderID, Message, and Type. It ensures that all necessary information is
	// captured and stored correctly when a notification is created.
	// Usage: When an order status changes or a delivery update occurs and a notification needs to be sent.
	// Function: Creates a new notification record and initializes its status.
	CreateNotification(ctx contextx.Contextx, notification *model.Notification) error

	// UpdateNotificationStatus updates the status of an existing notification. This method
	// changes the current status of the notification (e.g., from pending to sent)
	// and updates the timestamp to reflect the change. It ensures that the notification
	// status is accurately tracked and updated in real-time.
	// Usage: When the notification status changes, such as from pending to sent.
	// Function: Updates the current status of the notification record and logs the timestamp.
	UpdateNotificationStatus(ctx contextx.Contextx, notificationID string, status string) error

	// GetNotification retrieves a notification by its ID. This method fetches the details
	// of a specific notification based on the unique identifier provided. It returns
	// the Notification entity containing all relevant information.
	// Usage: When you need to query the details of a specific notification record.
	// Function: Retrieves the notification record based on its unique identifier.
	GetNotification(ctx contextx.Contextx, notificationID string) (item *model.Notification, err error)

	// ListNotificationsByUser retrieves a list of notifications for a specific user. This method
	// returns a paginated list of notifications for a given user, allowing for easier management
	// and tracking of multiple notifications.
	// Usage: When you need to manage and track notifications for a specific user.
	// Function: Provides a list of notifications for the specified user with pagination options.
	ListNotificationsByUser(
		ctx contextx.Contextx,
		userID string,
		options ListNotificationsOptions,
	) (items []*model.Notification, total int, err error)
}
