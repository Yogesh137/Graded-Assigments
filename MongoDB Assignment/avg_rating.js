db.products.aggregate([
    {
      $unwind: "$ratings"
    },
    {
      $group: {
        _id: "$productId",
        name: { $first: "$name" },
        averageRating: { $avg: "$ratings.rating" }
      }
    },
    {
      $match: { averageRating: { $gte: 4 } }
    },
    {
      $project: {
        _id: 0,
        productId: "$_id",
        name: 1,
        averageRating: 1
      }
    }
  ]);
  