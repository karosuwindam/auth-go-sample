import { Link } from "react-router-dom";

function page() {
    return (
        <div className="App">
            <header className="App-header">
                <p>
                    Test Page
                </p>
                <Link to="/" className='App-link'>Home</Link>
            </header>
        </div>
    )
}

export default page;