import Main from './pages/main-page';
import { Route, Routes, BrowserRouter as Router } from 'react-router-dom';
import UserPage from './pages/user-page';
import PrivateRoute from './authentication/private-route';
import { AuthProvider } from './authentication/auth-context';
import SignupPage from './pages/signup-page';
import RedirectUserRoute from './authentication/redirect-user-route';

function App() {
  return (
    <AuthProvider>
      <Router>
        <Routes>
          <Route path="/" element={<RedirectUserRoute />}>
            <Route path="/" element={<Main />} />
            <Route path="/signup" element={<SignupPage />} />
          </Route>
          <Route path="/" element={<PrivateRoute />}>
            <Route path="/userpage" element={<UserPage />} />
          </Route>
        </Routes>
      </Router>
    </AuthProvider>
  );
}

export default App;
