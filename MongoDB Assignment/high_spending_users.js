db.users.aggregate([
    {
      $lookup: {
        from: "orders",
        localField: "userId",
        foreignField: "userId",
        as: "userOrders"
      }
    },
    {
      $addFields: {
        totalSpent: {
          $sum: {
            $map: {
              input: "$userOrders",
              as: "order",
              in: "$$order.totalAmount"
            }
          }
        }
      }
    },
    {
      $match: { totalSpent: { $gt: 500 } }
    },
    {
      $project: {
        _id: 0,
        userId: 1,
        name: 1,
        totalSpent: 1
      }
    }
  ]);
  