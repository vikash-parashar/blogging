// main.js

// Function to fetch and display all blog posts
function displayAllPosts() {
    // Implement logic to send a GET request to fetch all posts
    // Use the JWT token from localStorage for authentication
    // Display the posts in the #blog-posts section
}

// Function to create a new post
function createPost() {
    // Implement logic to send a POST request to create a new post
    // Use the JWT token from localStorage for authentication
    // Handle success and error responses
}

// Function to delete a post by ID
function deletePost(postId) {
    // Implement logic to send a DELETE request to delete the post
    // Use the JWT token from localStorage for authentication
    // Handle success and error responses
}

// Function to edit a post by ID
function editPost(postId) {
    // Implement logic to send a PUT request to edit the post
    // Use the JWT token from localStorage for authentication
    // Handle success and error responses
}
// main.js

// Event listener for the "Delete Post" button
document.getElementById("delete-post-button").addEventListener("click", function () {
    const postId = // Get the ID of the post you want to delete (you can implement this logic)
    deletePost(postId);
});

// Event listener for the "Edit Post" button
document.getElementById("edit-post-button").addEventListener("click", function () {
    const postId = // Get the ID of the post you want to edit (you can implement this logic)
    editPost(postId);
});

// Event listeners for buttons
document.getElementById("display-posts-button").addEventListener("click", function () {
    displayAllPosts();
});

document.getElementById("create-post-button").addEventListener("click", function () {
    createPost();
});

// Initial setup
// Fetch and display posts when the page loads
displayAllPosts();
