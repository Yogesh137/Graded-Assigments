db.orders.aggregate([
    {
      $match: {
        orderDate: {
          $gte: new Date("2024-12-01T00:00:00Z"),
          $lte: new Date("2024-12-31T23:59:59Z")
        }
      }
    },
    {
      $lookup: {
        from: "users",
        localField: "userId",
        foreignField: "userId",
        as: "userDetails"
      }
    },
    {
      $unwind: "$userDetails"
    },
    {
      $project: {
        _id: 0,
        orderId: 1,
        orderDate: 1,
        totalAmount: 1,
        status: 1,
        userName: "$userDetails.name"
      }
    }
  ]);
  