package core

// A descriptor is a string representing the type of a field or method.
// Descriptors are represented in the class file format using modified UTF-8 strings (§4.4.7)
// and thus may be drawn, where not further constrained, from the entire Unicode codespace.
/**

FieldType term	Type	Interpretation
B	byte	signed byte
C	char	Unicode character code point in the Basic Multilingual Plane, encoded with UTF-16
D	double	double-precision floating-point value
F	float	single-precision floating-point value
I	int	integer
J	long	long integer
L ClassName ;	reference	an instance of class ClassName
S	short	signed short
Z	boolean	true or false
[	reference	one array dimension
*/
