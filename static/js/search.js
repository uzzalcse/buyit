
document.addEventListener('DOMContentLoaded', function () {
    const searchInput = document.getElementById('search-input');
    const searchResults = document.getElementById('search-results');
    const productDetails = document.getElementById('product-details');
    let debounceTimer;

    // Initially hide both results and details
    searchResults.classList.remove('active');
    productDetails.classList.remove('active');

    searchInput.addEventListener('input', function (e) {
        clearTimeout(debounceTimer);
        debounceTimer = setTimeout(() => {
            const query = e.target.value.trim();
            if (query.length >= 2) {
                searchProducts(query);
                searchResults.classList.add('active');
            } else {
                searchResults.innerHTML = '';
                searchResults.classList.remove('active');
            }
        }, 300);
    });

    // Focus event to show results if there's input
    searchInput.addEventListener('focus', function (e) {
        if (this.value.trim().length >= 2) {
            searchResults.classList.add('active');
        }
    });

    // Close search results when clicking outside
    document.addEventListener('click', function (e) {
        if (!searchResults.contains(e.target) && 
            !searchInput.contains(e.target)) {
            searchResults.classList.remove('active');
        }
    });

    async function searchProducts(query) {
        try {
            const response = await fetch(`/api/products/search?q=${encodeURIComponent(query)}`);
            const data = await response.json();

            if (!data.hits?.hits?.length) {
                searchResults.innerHTML = '<p class="product-item">No products found</p>';
                return;
            }

            const productsList = data.hits.hits.flatMap(hit => {
                const orderId = hit._id;
                return (hit._source.products || []).map(product => ({
                    ...product,
                    orderId: orderId
                }));
            });

            if (productsList.length === 0) {
                searchResults.innerHTML = '<p class="product-item">No products found</p>';
                return;
            }

            searchResults.innerHTML = productsList.map(product => `
                <div class="product-item" 
                     data-product-id="${product.product_id}"
                     data-order-id="${product.orderId}">
                    ${product.product_name}
                </div>
            `).join('');

        } catch (error) {
            console.error('Error searching products:', error);
            searchResults.innerHTML = '<p class="product-item">Error searching products</p>';
        }
    }

    searchResults.addEventListener('click', function (e) {
        const productItem = e.target.closest('.product-item');
        if (productItem) {
            const productId = productItem.dataset.productId;
            const orderId = productItem.dataset.orderId;
            getProductDetails(orderId, productId);
            searchResults.classList.remove('active');
        }
    });

    async function getProductDetails(orderId, productId) {
        try {
            const response = await fetch(`/api/products/${orderId}`);
            const data = await response.json();

            if (!data || !data._source || !data._source.products) {
                productDetails.innerHTML = '<p>Product details not found</p>';
                return;
            }

            const product = data._source.products.find(p => p.product_id.toString() === productId.toString());
            
            if (!product) {
                productDetails.innerHTML = '<p>Product details not found</p>';
                return;
            }

            productDetails.innerHTML = `
                <div class="product-details">
                    <h2>${product.product_name}</h2>
                    <p>Category: ${product.category}</p>
                    <p>Price: €${product.price}</p>
                    <p>Manufacturer: ${product.manufacturer}</p>
                    <p>SKU: ${product.sku}</p>
                    <p>Order ID: ${orderId}</p>
                </div>
            `;
            productDetails.classList.add('active');

        } catch (error) {
            console.error('Error fetching product details:', error);
            productDetails.innerHTML = '<p>Error loading product details</p>';
        }
    }
});
