"""Hello unit test module."""

from py_observability_kit.hello import hello


def test_hello():
    """Test the hello function."""
    assert hello() == "Hello py-observability-kit"
