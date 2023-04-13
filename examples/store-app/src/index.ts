import {DhClient} from "data-handler-client";
import {Order, Product} from "./schema";

const client = new DhClient({
    Addr: "127.0.0.1:9009",
    Insecure: true,
})

async function run() {
    await client.authenticateWithUsernameAndPassword("admin", "admin")
    const productRepo = client.newRepository<Product>("default", "product")
    const orderRepo = client.newRepository<Order>("default", "order")

    const extension = client.NewExtensionService("127.0.0.1", 17686)

    await extension.run()

    const orderExtension = orderRepo.extend(extension)

    orderExtension.onCreate(async (order) => {
        if (order.status != 'pending') {
            throw new Error('Order must be created with pending status')
        }

        const product = await productRepo.get(order.product.id)

        if (product.quantity < order.quantity) {
            throw new Error('Not enough product in stock')
        }

        return order
    })

    orderExtension.onUpdate(async (order) => {
        const existingOrder = await orderRepo.get(order.id)

        if (existingOrder.status == 'completed') {
            throw new Error('Cannot update completed order')
        }

        if (order.status == 'completed') {
            const product = await productRepo.get(order.product.id)

            if (product.quantity < order.quantity) {
                throw new Error('Not enough product in stock')
            }

            product.quantity -= order.quantity
            await productRepo.update(product)
        }

        return order
    })
}

run().then()