import { Button, Navbar } from "flowbite-react";
import { useRouter } from "next/router";
import logo from '../assets/logo.png';

export default function Layout(props: { children: React.ReactNode }) {
    const router = useRouter();
    return (
        <>
            <Navbar
                fluid={true}
                rounded={true}
            >
                <Navbar.Brand href="/">
                    <img src={logo.src} className="mr-3 h-6 sm:h-9"alt="Flowbite Logo"/>
                    <span className="self-center whitespace-nowrap text-xl font-semibold dark:text-white">
                        Face-To-Face
                    </span>
                </Navbar.Brand>
                <div className="flex md:order-2 mr-0 ml-auto md:hidden">
                    <Button onClick={() => router.push('/login')}>
                        Get Started
                    </Button>
                </div>
                <Navbar.Toggle />
                <Navbar.Collapse className="mr-5 ml-auto" >
                    <Navbar.Link
                        href="/"
                        active={true}
                    >
                        Home
                    </Navbar.Link>
                    <Navbar.Link href="/navbars">
                        About
                    </Navbar.Link>
                    <Navbar.Link href="/navbars">
                        Pricing
                    </Navbar.Link>
                </Navbar.Collapse>
                <Button className="hidden md:block" onClick={() => router.push('/login')}>
                    Get Started
                </Button>
            </Navbar>
            <main>{props.children}</main>
        </>
    )
} 