# adcl Generator

This package contains the adcl generator for data structures and functions
corresponding to known ADC commands. The following is generated for each known
ADC command (as well as utility types and functions):

 * `message.???Content` - struct type, written to `message/content_???.go`
 * `parser.Parse???Content()` - function, written to `parser/parse_???.go`
 * `builder.Build???Content()` - function, written to `builder/build_???.go`

## Concepts and Message Model

 * Messages correspond to the ADC message schema associated with each command
   and result in a `???Content` struct type being generated. They are specified
   in the `messages.md` file.
 * Logical parameters of messages are specified along messages in the
   `messages.md` file.
     * Each logical parameter has a logical type.
     * Each logical parameter is either a positional or a named parameter.
     * Each logical parameter results in a field in the messages struct type.
     * Each logical parameter also results in a string field `str...` in the
       messages struct type for storing the raw parameter value.

## Mapper

Mappers are responsible for generating code which handles ADC related
functionality for specific logical parameters. The generator chooses the mapper
of a logical parameter based on the parameter's logical type. Alternatively, the
logical parameter may override the type's default mapper and specify a different
one.
