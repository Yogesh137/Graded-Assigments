db.orders.find({ orderId: "ORD001" }).forEach(order => {
    order.items.forEach(item => {
      db.products.updateOne(
        { productId: item.productId },
        { $inc: { stock: -item.quantity } }
      );
    });
  });
  