package redfish

import "goprom/internal/models"

func (c *Client) GetTemperatures(endpoints []string) ([]*models.Temperature, error) {
	var results []*models.Temperature

	for _, endpoint := range endpoints {
		var temp models.Temperature

		status, err := c.Get(endpoint, &temp)
		if err != nil {
			continue
		}

		if status == 404 {
			continue
		}

		results = append(results, &temp)
	}

	return results, nil
}
