//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

// ILogisticsHandler defines the interface for logistics handler operations.
type ILogisticsHandler interface {
}
