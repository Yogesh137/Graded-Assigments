db.warehouses.aggregate([
    {
      $geoNear: {
        near: { type: "Point", coordinates: [-74.006, 40.7128] },
        distanceField: "distance",
        maxDistance: 50000, // 50 kilometers in meters
        spherical: true,
        query: { products: "P001" }
      }
    },
    {
      $project: {
        _id: 0,
        warehouseId: 1,
        distance: 1,
        location: 1,
        products: 1
      }
    }
  ]);
  