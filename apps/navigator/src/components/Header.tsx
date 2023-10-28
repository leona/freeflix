import { Link } from 'preact-router/match';

export function Header() {
	return (
		<header>
			<nav>
				<Link href="/" activeClassName="active">
					Search
				</Link>
				<Link href="/downloads" activeClassName="active">
					Downloads
				</Link>
			</nav>
		</header>
	);
}
