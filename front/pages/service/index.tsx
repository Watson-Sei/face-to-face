import useRequireLogin from "../../hooks/useRequireLogin";

export default function Index() {

    const { loading } = useRequireLogin("guest");

    if (loading) return <div>ローディング中</div>

    return (
        <div>This is Service Main Page</div>
    )
}