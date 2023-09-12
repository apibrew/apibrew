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

public class `)
//line templates/java/resource.qtpl:10
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:10
	qw422016.N().S(` extends Entity {
`)
//line templates/java/resource.qtpl:11
	for _, property := range resource.Properties {
//line templates/java/resource.qtpl:11
		qw422016.N().S(`    private `)
//line templates/java/resource.qtpl:12
		qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:12
		qw422016.N().S(` `)
//line templates/java/resource.qtpl:12
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:12
		qw422016.N().S(`;
`)
//line templates/java/resource.qtpl:13
	}
//line templates/java/resource.qtpl:13
	qw422016.N().S(`
    public static final EntityInfo<`)
//line templates/java/resource.qtpl:15
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:15
	qw422016.N().S(`> entityInfo = new EntityInfo<>("`)
//line templates/java/resource.qtpl:15
	qw422016.E().S(resource.Namespace)
//line templates/java/resource.qtpl:15
	qw422016.N().S(`", "`)
//line templates/java/resource.qtpl:15
	qw422016.E().S(resource.Name)
//line templates/java/resource.qtpl:15
	qw422016.N().S(`", `)
//line templates/java/resource.qtpl:15
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:15
	qw422016.N().S(`.class);

`)
//line templates/java/resource.qtpl:17
	for _, subType := range resource.Types {
//line templates/java/resource.qtpl:17
		qw422016.N().S(`    public static class `)
//line templates/java/resource.qtpl:18
		qw422016.E().S(javaClassName(subType.Name))
//line templates/java/resource.qtpl:18
		qw422016.N().S(` {
`)
//line templates/java/resource.qtpl:19
		for _, property := range subType.Properties {
//line templates/java/resource.qtpl:19
			qw422016.N().S(`        private `)
//line templates/java/resource.qtpl:20
			qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:20
			qw422016.N().S(` `)
//line templates/java/resource.qtpl:20
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:20
			qw422016.N().S(`;
`)
//line templates/java/resource.qtpl:21
		}
//line templates/java/resource.qtpl:21
		qw422016.N().S(`
`)
//line templates/java/resource.qtpl:23
		for _, property := range subType.Properties {
//line templates/java/resource.qtpl:23
			qw422016.N().S(`        public `)
//line templates/java/resource.qtpl:24
			qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:24
			qw422016.N().S(` get`)
//line templates/java/resource.qtpl:24
			qw422016.N().S(javaClassName(property.Name))
//line templates/java/resource.qtpl:24
			qw422016.N().S(`() {
            return `)
//line templates/java/resource.qtpl:25
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:25
			qw422016.N().S(`;
        }

        public void set`)
//line templates/java/resource.qtpl:28
			qw422016.N().S(javaClassName(property.Name))
//line templates/java/resource.qtpl:28
			qw422016.N().S(`(`)
//line templates/java/resource.qtpl:28
			qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:28
			qw422016.N().S(` `)
//line templates/java/resource.qtpl:28
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:28
			qw422016.N().S(`) {
            this.`)
//line templates/java/resource.qtpl:29
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:29
			qw422016.N().S(` = `)
//line templates/java/resource.qtpl:29
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:29
			qw422016.N().S(`;
        }
`)
//line templates/java/resource.qtpl:31
		}
//line templates/java/resource.qtpl:31
		qw422016.N().S(`    }
`)
//line templates/java/resource.qtpl:33
	}
//line templates/java/resource.qtpl:33
	qw422016.N().S(`
`)
//line templates/java/resource.qtpl:35
	for _, enum := range getAllEnums(resource) {
//line templates/java/resource.qtpl:35
		qw422016.N().S(`    public static enum `)
//line templates/java/resource.qtpl:36
		qw422016.E().S(javaClassName(enum.Name))
//line templates/java/resource.qtpl:36
		qw422016.N().S(` {
`)
//line templates/java/resource.qtpl:37
		for index, enumValue := range enum.EnumValues {
//line templates/java/resource.qtpl:37
			qw422016.N().S(`        `)
//line templates/java/resource.qtpl:38
			qw422016.N().S(enumName(enumValue))
//line templates/java/resource.qtpl:38
			qw422016.N().S(`("`)
//line templates/java/resource.qtpl:38
			qw422016.E().S(enumValue)
//line templates/java/resource.qtpl:38
			qw422016.N().S(`")`)
//line templates/java/resource.qtpl:38
			if index < len(enum.EnumValues)-1 {
//line templates/java/resource.qtpl:38
				qw422016.N().S(`,`)
//line templates/java/resource.qtpl:38
			} else {
//line templates/java/resource.qtpl:38
				qw422016.N().S(`;`)
//line templates/java/resource.qtpl:38
			}
//line templates/java/resource.qtpl:38
			qw422016.N().S(`
`)
//line templates/java/resource.qtpl:39
		}
//line templates/java/resource.qtpl:39
		qw422016.N().S(`
        private final String value;

        `)
//line templates/java/resource.qtpl:43
		qw422016.E().S(javaClassName(enum.Name))
//line templates/java/resource.qtpl:43
		qw422016.N().S(`(String value) {
            this.value = value;
        }

        public String getValue() {
            return value;
        }
    }
`)
//line templates/java/resource.qtpl:51
	}
//line templates/java/resource.qtpl:51
	qw422016.N().S(`
    public `)
//line templates/java/resource.qtpl:53
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:53
	qw422016.N().S(`() {
    }

    public EntityInfo<`)
//line templates/java/resource.qtpl:56
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:56
	qw422016.N().S(`> getEntityInfo() {
        return entityInfo;
    }

`)
//line templates/java/resource.qtpl:60
	for _, property := range resource.Properties {
//line templates/java/resource.qtpl:60
		qw422016.N().S(`    public `)
//line templates/java/resource.qtpl:61
		qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:61
		qw422016.N().S(` get`)
//line templates/java/resource.qtpl:61
		qw422016.N().S(javaClassName(property.Name))
//line templates/java/resource.qtpl:61
		qw422016.N().S(`() {
        return `)
//line templates/java/resource.qtpl:62
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:62
		qw422016.N().S(`;
    }

    public void set`)
//line templates/java/resource.qtpl:65
		qw422016.N().S(javaClassName(property.Name))
//line templates/java/resource.qtpl:65
		qw422016.N().S(`(`)
//line templates/java/resource.qtpl:65
		qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:65
		qw422016.N().S(` `)
//line templates/java/resource.qtpl:65
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:65
		qw422016.N().S(`) {
        this.`)
//line templates/java/resource.qtpl:66
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:66
		qw422016.N().S(` = `)
//line templates/java/resource.qtpl:66
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:66
		qw422016.N().S(`;
    }
`)
//line templates/java/resource.qtpl:68
	}
//line templates/java/resource.qtpl:68
	qw422016.N().S(`
    @Override
    public boolean equals(Object o) {
        if (!(o instanceof `)
//line templates/java/resource.qtpl:72
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:72
	qw422016.N().S(`)) {
            return false;
        }

        `)
//line templates/java/resource.qtpl:76
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:76
	qw422016.N().S(` obj = (`)
//line templates/java/resource.qtpl:76
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:76
	qw422016.N().S(`) o;

`)
//line templates/java/resource.qtpl:78
	for _, property := range resource.Properties {
//line templates/java/resource.qtpl:78
		qw422016.N().S(`        if (!Objects.equals(this.`)
//line templates/java/resource.qtpl:79
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:79
		qw422016.N().S(`, obj.`)
//line templates/java/resource.qtpl:79
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:79
		qw422016.N().S(`)) {
            return false;
        }
`)
//line templates/java/resource.qtpl:82
	}
//line templates/java/resource.qtpl:82
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
//line templates/java/resource.qtpl:98
}

//line templates/java/resource.qtpl:98
func WriteGenerateClassCode(qq422016 qtio422016.Writer, pkg string, resource *model.Resource) {
//line templates/java/resource.qtpl:98
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/java/resource.qtpl:98
	StreamGenerateClassCode(qw422016, pkg, resource)
//line templates/java/resource.qtpl:98
	qt422016.ReleaseWriter(qw422016)
//line templates/java/resource.qtpl:98
}

//line templates/java/resource.qtpl:98
func GenerateClassCode(pkg string, resource *model.Resource) string {
//line templates/java/resource.qtpl:98
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/java/resource.qtpl:98
	WriteGenerateClassCode(qb422016, pkg, resource)
//line templates/java/resource.qtpl:98
	qs422016 := string(qb422016.B)
//line templates/java/resource.qtpl:98
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/java/resource.qtpl:98
	return qs422016
//line templates/java/resource.qtpl:98
}
