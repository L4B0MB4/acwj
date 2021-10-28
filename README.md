# acwj

Writing a compiler in go.

I use this to get more familiar with go and to learn more about compilers :)

I'm following the journey from: https://github.com/DoctorWkt/acwj

Starting at part 12 I started going in my own direction

# Concept

### Node:

```

node.{id|Root}{

    state: { //n-times
        ...
        {id}: basictype|node|basictype[]|node[]
    }

    connection: node

    persist: nodefunc

    load: nodefunc

    changeState: nodefunc

    _parent:node

    _children:node[]

    {//Rootonly

        _onTrigger:helperfunc
    }
}

```

### NodeFunc:

```
    nodeid.{load|persist|changeState} {

        //usual code and helper function calls

    }

```

### HelperFunc:

```
    common.{id}(params) {

        //usual code and helper function calls

    }

```

## Example:

```
    node.Root{

        state={
            counter:0
        }

        load: ()=>{
            counter = 1
        }

        changeState : ()=>{
            counter+=1
        }

        onTrigger: (params)=>{

        }

    }

```
