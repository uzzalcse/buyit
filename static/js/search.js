document.addEventListener('DOMContentLoaded', function() {
    const searchInput = document.getElementById('search-input');
    const resultsContainer = document.getElementById('search-results');
    let debounceTimer;

    searchInput.addEventListener('input', function(e) {
        clearTimeout(debounceTimer);
        debounceTimer = setTimeout(() => {
            const query = e.target.value.trim();
            if (query.length >= 2) {
                fetchResults(query);
            } else {
                resultsContainer.innerHTML = '';
            }
        }, 300);
    });

    async function fetchResults(query) {
        try {
            const response = await fetch(`/api/products/search?q=${encodeURIComponent(query)}`);
            const data = await response.json();

            // Extracting products from each hit
            const products = data.hits.hits.flatMap(hit => hit._source.products || []);

            if (products.length === 0) {
                resultsContainer.innerHTML = '<p>No products found</p>';
                return;
            }

            // Display results
            resultsContainer.innerHTML = products.map(product => `
                <div class="product-item">
                    <h3>${product.product_name}</h3>
                    <p>Category: ${product.category}</p>
                    <p>Price: €${product.price}</p>
                    <p>Manufacturer: ${product.manufacturer}</p>
                </div>
            `).join('');

        } catch (error) {
            console.error('Error fetching results:', error);
            resultsContainer.innerHTML = '<p>Error fetching results</p>';
        }
    }
});
