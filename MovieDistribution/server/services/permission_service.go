package services

import "MovieDistribution/models"

func HasAccess(d *models.Distributor, region *models.Region) bool {
	r := region

	for r != nil {
		if d.Includes[r.Code] {
			if d.Parent != nil {
				return HasAccess(d.Parent, region)
			}
			return true
		}

		if d.Excludes[r.Code] {
			return false
		}

		r = r.Parent
	}

	return false
}
