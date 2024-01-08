import AppHeader from "../components/Header";
import WorkspaceBar from "../components/WorkspaceBar";

function DashboardPage() {
    return (
        <>
            <AppHeader />
            <div className="workspace">
                <WorkspaceBar />
            </div>
        </>
    );
}

export default DashboardPage;