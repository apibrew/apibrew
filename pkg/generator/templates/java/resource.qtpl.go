// Code generated by qtc from "resource.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/java/resource.qtpl:1
package java

//line templates/java/resource.qtpl:1
import "github.com/apibrew/apibrew/pkg/model"

//line templates/java/resource.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/java/resource.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/java/resource.qtpl:3
func StreamGenerateClassCode(qw422016 *qt422016.Writer, pkg string, resource *model.Resource) {
//line templates/java/resource.qtpl:3
	qw422016.N().S(`package `)
//line templates/java/resource.qtpl:4
	qw422016.E().S(pkg)
//line templates/java/resource.qtpl:4
	qw422016.N().S(`;

import java.util.Objects;
import io.apibrew.lib.EntityInfo;
import io.apibrew.lib.Entity;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class `)
//line templates/java/resource.qtpl:15
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:15
	qw422016.N().S(` extends Entity {
`)
//line templates/java/resource.qtpl:16
	for _, property := range resource.Properties {
//line templates/java/resource.qtpl:16
		qw422016.N().S(`    `)
//line templates/java/resource.qtpl:17
		qw422016.N().S(getJavaPropertyAnnotations(resource, property))
//line templates/java/resource.qtpl:17
		qw422016.N().S(`
    private `)
//line templates/java/resource.qtpl:18
		qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:18
		qw422016.N().S(` `)
//line templates/java/resource.qtpl:18
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:18
		qw422016.N().S(`;
`)
//line templates/java/resource.qtpl:19
	}
//line templates/java/resource.qtpl:19
	qw422016.N().S(`
    @JsonIgnore
    public static final EntityInfo<`)
//line templates/java/resource.qtpl:22
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:22
	qw422016.N().S(`> entityInfo = new EntityInfo<>("`)
//line templates/java/resource.qtpl:22
	qw422016.E().S(resource.Namespace)
//line templates/java/resource.qtpl:22
	qw422016.N().S(`", "`)
//line templates/java/resource.qtpl:22
	qw422016.E().S(resource.Name)
//line templates/java/resource.qtpl:22
	qw422016.N().S(`", `)
//line templates/java/resource.qtpl:22
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:22
	qw422016.N().S(`.class, "`)
//line templates/java/resource.qtpl:22
	qw422016.E().S(getRestPath(resource))
//line templates/java/resource.qtpl:22
	qw422016.N().S(`");

`)
//line templates/java/resource.qtpl:24
	for _, subType := range resource.Types {
//line templates/java/resource.qtpl:24
		qw422016.N().S(`    public static class `)
//line templates/java/resource.qtpl:25
		qw422016.E().S(javaClassName(subType.Name))
//line templates/java/resource.qtpl:25
		qw422016.N().S(` {
`)
//line templates/java/resource.qtpl:26
		for _, property := range subType.Properties {
//line templates/java/resource.qtpl:26
			qw422016.N().S(`        `)
//line templates/java/resource.qtpl:27
			qw422016.N().S(getJavaPropertyAnnotations(resource, property))
//line templates/java/resource.qtpl:27
			qw422016.N().S(`
        private `)
//line templates/java/resource.qtpl:28
			qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:28
			qw422016.N().S(` `)
//line templates/java/resource.qtpl:28
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:28
			qw422016.N().S(`;
`)
//line templates/java/resource.qtpl:29
		}
//line templates/java/resource.qtpl:29
		qw422016.N().S(`
`)
//line templates/java/resource.qtpl:31
		for _, property := range subType.Properties {
//line templates/java/resource.qtpl:31
			qw422016.N().S(`        public `)
//line templates/java/resource.qtpl:32
			qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:32
			qw422016.N().S(` get`)
//line templates/java/resource.qtpl:32
			qw422016.N().S(javaClassName(property.Name))
//line templates/java/resource.qtpl:32
			qw422016.N().S(`() {
            return `)
//line templates/java/resource.qtpl:33
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:33
			qw422016.N().S(`;
        }

        public void set`)
//line templates/java/resource.qtpl:36
			qw422016.N().S(javaClassName(property.Name))
//line templates/java/resource.qtpl:36
			qw422016.N().S(`(`)
//line templates/java/resource.qtpl:36
			qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:36
			qw422016.N().S(` `)
//line templates/java/resource.qtpl:36
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:36
			qw422016.N().S(`) {
            this.`)
//line templates/java/resource.qtpl:37
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:37
			qw422016.N().S(` = `)
//line templates/java/resource.qtpl:37
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:37
			qw422016.N().S(`;
        }

        public `)
//line templates/java/resource.qtpl:40
			qw422016.E().S(javaClassName(subType.Name))
//line templates/java/resource.qtpl:40
			qw422016.N().S(` with`)
//line templates/java/resource.qtpl:40
			qw422016.N().S(javaClassName(property.Name))
//line templates/java/resource.qtpl:40
			qw422016.N().S(`(`)
//line templates/java/resource.qtpl:40
			qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:40
			qw422016.N().S(` `)
//line templates/java/resource.qtpl:40
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:40
			qw422016.N().S(`) {
            this.`)
//line templates/java/resource.qtpl:41
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:41
			qw422016.N().S(` = `)
//line templates/java/resource.qtpl:41
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:41
			qw422016.N().S(`;

            return this;
        }
`)
//line templates/java/resource.qtpl:45
		}
//line templates/java/resource.qtpl:45
		qw422016.N().S(`    }
`)
//line templates/java/resource.qtpl:47
	}
//line templates/java/resource.qtpl:47
	qw422016.N().S(`
`)
//line templates/java/resource.qtpl:49
	for _, enum := range getAllEnums(resource) {
//line templates/java/resource.qtpl:49
		qw422016.N().S(`    public static enum `)
//line templates/java/resource.qtpl:50
		qw422016.E().S(javaClassName(enum.Name))
//line templates/java/resource.qtpl:50
		qw422016.N().S(` {
`)
//line templates/java/resource.qtpl:51
		for index, enumValue := range enum.EnumValues {
//line templates/java/resource.qtpl:51
			qw422016.N().S(`        `)
//line templates/java/resource.qtpl:52
			qw422016.N().S(enumName(enumValue))
//line templates/java/resource.qtpl:52
			qw422016.N().S(`("`)
//line templates/java/resource.qtpl:52
			qw422016.E().S(enumValue)
//line templates/java/resource.qtpl:52
			qw422016.N().S(`")`)
//line templates/java/resource.qtpl:52
			if index < len(enum.EnumValues)-1 {
//line templates/java/resource.qtpl:52
				qw422016.N().S(`,`)
//line templates/java/resource.qtpl:52
			} else {
//line templates/java/resource.qtpl:52
				qw422016.N().S(`;`)
//line templates/java/resource.qtpl:52
			}
//line templates/java/resource.qtpl:52
			qw422016.N().S(`
`)
//line templates/java/resource.qtpl:53
		}
//line templates/java/resource.qtpl:53
		qw422016.N().S(`
        private final String value;

        `)
//line templates/java/resource.qtpl:57
		qw422016.E().S(javaClassName(enum.Name))
//line templates/java/resource.qtpl:57
		qw422016.N().S(`(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }
`)
//line templates/java/resource.qtpl:66
	}
//line templates/java/resource.qtpl:66
	qw422016.N().S(`
    public `)
//line templates/java/resource.qtpl:68
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:68
	qw422016.N().S(`() {
    }

`)
//line templates/java/resource.qtpl:71
	for _, property := range resource.Properties {
//line templates/java/resource.qtpl:71
		qw422016.N().S(`    public `)
//line templates/java/resource.qtpl:72
		qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:72
		qw422016.N().S(` get`)
//line templates/java/resource.qtpl:72
		qw422016.N().S(javaClassName(property.Name))
//line templates/java/resource.qtpl:72
		qw422016.N().S(`() {
        return `)
//line templates/java/resource.qtpl:73
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:73
		qw422016.N().S(`;
    }

    public void set`)
//line templates/java/resource.qtpl:76
		qw422016.N().S(javaClassName(property.Name))
//line templates/java/resource.qtpl:76
		qw422016.N().S(`(`)
//line templates/java/resource.qtpl:76
		qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:76
		qw422016.N().S(` `)
//line templates/java/resource.qtpl:76
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:76
		qw422016.N().S(`) {
        this.`)
//line templates/java/resource.qtpl:77
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:77
		qw422016.N().S(` = `)
//line templates/java/resource.qtpl:77
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:77
		qw422016.N().S(`;
    }

    public `)
//line templates/java/resource.qtpl:80
		qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:80
		qw422016.N().S(` with`)
//line templates/java/resource.qtpl:80
		qw422016.N().S(javaClassName(property.Name))
//line templates/java/resource.qtpl:80
		qw422016.N().S(`(`)
//line templates/java/resource.qtpl:80
		qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:80
		qw422016.N().S(` `)
//line templates/java/resource.qtpl:80
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:80
		qw422016.N().S(`) {
        this.`)
//line templates/java/resource.qtpl:81
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:81
		qw422016.N().S(` = `)
//line templates/java/resource.qtpl:81
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:81
		qw422016.N().S(`;

        return this;
    }
`)
//line templates/java/resource.qtpl:85
	}
//line templates/java/resource.qtpl:85
	qw422016.N().S(`
    @Override
    public boolean equals(Object o) {
        if (!(o instanceof `)
//line templates/java/resource.qtpl:89
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:89
	qw422016.N().S(`)) {
            return false;
        }

        `)
//line templates/java/resource.qtpl:93
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:93
	qw422016.N().S(` obj = (`)
//line templates/java/resource.qtpl:93
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:93
	qw422016.N().S(`) o;

`)
//line templates/java/resource.qtpl:95
	for _, property := range resource.Properties {
//line templates/java/resource.qtpl:95
		qw422016.N().S(`        if (!Objects.equals(this.`)
//line templates/java/resource.qtpl:96
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:96
		qw422016.N().S(`, obj.`)
//line templates/java/resource.qtpl:96
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:96
		qw422016.N().S(`)) {
            return false;
        }
`)
//line templates/java/resource.qtpl:99
	}
//line templates/java/resource.qtpl:99
	qw422016.N().S(`
        return true;
    }

    @Override
    public int hashCode() {
        if (id == null) {
            return super.hashCode();
        }

        return id.hashCode();
    }
}


`)
//line templates/java/resource.qtpl:115
}

//line templates/java/resource.qtpl:115
func WriteGenerateClassCode(qq422016 qtio422016.Writer, pkg string, resource *model.Resource) {
//line templates/java/resource.qtpl:115
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/java/resource.qtpl:115
	StreamGenerateClassCode(qw422016, pkg, resource)
//line templates/java/resource.qtpl:115
	qt422016.ReleaseWriter(qw422016)
//line templates/java/resource.qtpl:115
}

//line templates/java/resource.qtpl:115
func GenerateClassCode(pkg string, resource *model.Resource) string {
//line templates/java/resource.qtpl:115
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/java/resource.qtpl:115
	WriteGenerateClassCode(qb422016, pkg, resource)
//line templates/java/resource.qtpl:115
	qs422016 := string(qb422016.B)
//line templates/java/resource.qtpl:115
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/java/resource.qtpl:115
	return qs422016
//line templates/java/resource.qtpl:115
}
