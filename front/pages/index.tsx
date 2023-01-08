import useRequireLogin from '../hooks/useRequireLogin'

export default function LP() {
  const { loading } = useRequireLogin("guest");

  if (loading) return <div>ローディング中</div>

  return (
    <div>
      <h1 className='text-3xl font-bold underline'>This is LP page.</h1>
    </div>
  )
}