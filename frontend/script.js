const form = document.querySelector('.glass-card');
const toastContainer = document.getElementById('toast-container');

form.addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const nameInput = document.getElementById('username');
    const submitBtn = form.querySelector('button');
    const memberName = nameInput.value;

    // Loading state
    submitBtn.innerText = "Verifying...";
    submitBtn.disabled = true;

    try {
        const response = await fetch(`/member?member=${encodeURIComponent(memberName)}`);
        const data = await response.text();

        if (response.ok) {
            showToast(data, 'success');
        } else {
            showToast(data || "Member not found", 'error');
        }
    } catch (error) {
        showToast("Connection to Edge Server failed", 'error');
    } finally {
        submitBtn.innerText = "Complete Setup";
        submitBtn.disabled = false;
    }
});

function showToast(message, type) {
    const toast = document.createElement('div');
    toast.className = `toast ${type}`;
    
    // Add an icon based on type
    const icon = type === 'success' ? '✓' : '✕';
    
    toast.innerHTML = `
        <span class="toast-icon">${icon}</span>
        <span class="toast-message">${message}</span>
    `;

    toastContainer.appendChild(toast);

    // Remove toast after 4 seconds
    setTimeout(() => {
        toast.classList.add('fade-out');
        toast.addEventListener('animationend', () => {
            toast.remove();
        });
    }, 4000);
}